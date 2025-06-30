// Copyright 2025 Albert Kapitanov.
// SPDX-License-Identifier: BSD-3-Clause
//
// This file is NOT a part of the original goimports tool from the Go project.

package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strconv"

	"golang.org/x/tools/internal/imports"
)

// The idea here is simple:
// Before running the goimports tool, we need to preprocess the file:
// each import spec is replaced with an equivalent one that does not have token position information.
// Here is an example of such a transformation:
//
//   import (
//
//     "fmt"
//
//     "local.example.com/group/project/shared"
//
//     "github.com/spf13/cobra"
//
//   )
//
//   import (
//     "fmt"
//     "local.example.com/group/project/shared"
//     "github.com/spf13/cobra"
//   )
//
// The first import spec has token positions, the second one does not - therefore we effectively trim empty lines caused by IDE auto-imports.
//
// Known limitations:
// - If the import spec has a comment, we cannot transform it, because the comment's position will be lost.
//   We have to leave it as is, skipping the transformation.

func preprocessFile(filename string, src []byte, opt *imports.Options) []byte {
	fset := token.NewFileSet()
	var parserMode parser.Mode
	if opt.Comments {
		parserMode |= parser.ParseComments
	}
	if opt.AllErrors {
		parserMode |= parser.AllErrors
	}

	file, err := parser.ParseFile(fset, filename, src, parserMode)
	if err != nil {
		return src
	}

	scanFileImports(file)

	printerMode := printer.UseSpaces
	if opt.TabIndent {
		printerMode |= printer.TabIndent
	}
	printConfig := &printer.Config{Mode: printerMode, Tabwidth: opt.TabWidth}
	var buf bytes.Buffer
	err = printConfig.Fprint(&buf, fset, file)
	if err != nil {
		return src
	}
	out := buf.Bytes()
	return out
}

func scanFileImports(file *ast.File) {
	for i := 0; i < len(file.Decls); i++ {
		decl := file.Decls[i]
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.IMPORT {
			continue
		}

		decl = transformGenDecl(gen)
		file.Decls[i] = decl
	}
}

func transformGenDecl(decl *ast.GenDecl) *ast.GenDecl {
	for _, spec := range decl.Specs {
		impspec := spec.(*ast.ImportSpec)
		if importPath(impspec) == "C" {
			return decl // It's a CGO import.
		}
	}

	result := &ast.GenDecl{
		Doc:    decl.Doc,
		Tok:    decl.Tok,
		TokPos: decl.TokPos,
		Lparen: decl.Lparen,
		Rparen: token.NoPos,
		Specs:  make([]ast.Spec, 0, len(decl.Specs)),
	}

	for _, spec := range decl.Specs {
		importSpec, ok := spec.(*ast.ImportSpec)
		if !ok {
			result.Specs = append(result.Specs, spec)
			continue
		}

		transformedImportSpec, ok := transformImportSpec(importSpec)
		if !ok {
			return decl
		}

		result.Specs = append(result.Specs, transformedImportSpec)
	}

	return result
}

func transformImportSpec(importSpec *ast.ImportSpec) (*ast.ImportSpec, bool) {
	hasDoc := importSpec.Doc != nil && len(importSpec.Doc.List) > 0
	hasComment := importSpec.Comment != nil && len(importSpec.Comment.List) > 0

	if hasDoc || hasComment {
		// If the import has a comment, we cannot proceed with the transformation;
		// otherwise, the comment's position will be lost.
		return nil, false
	}

	transformedImportSpec := &ast.ImportSpec{
		Doc:     transformCommentGroup(importSpec.Doc),
		Name:    transformIdent(importSpec.Name),
		Path:    transformBasicLit(importSpec.Path),
		Comment: transformCommentGroup(importSpec.Comment),
		EndPos:  token.NoPos,
	}

	return transformedImportSpec, true
}

func importPath(s ast.Spec) string {
	t, err := strconv.Unquote(s.(*ast.ImportSpec).Path.Value)
	if err == nil {
		return t
	}
	return ""
}

func transformCommentGroup(cg *ast.CommentGroup) *ast.CommentGroup {
	if cg == nil {
		return nil
	}
	transformed := &ast.CommentGroup{
		List: make([]*ast.Comment, len(cg.List)),
	}
	for i, comment := range cg.List {

		transformed.List[i] = &ast.Comment{
			Slash: token.NoPos,
			Text:  comment.Text,
		}
	}
	return transformed
}

func transformIdent(id *ast.Ident) *ast.Ident {
	if id == nil {
		return nil
	}
	return &ast.Ident{
		NamePos: token.NoPos,
		Name:    id.Name,
	}
}

func transformBasicLit(bl *ast.BasicLit) *ast.BasicLit {
	if bl == nil {
		return nil
	}
	return &ast.BasicLit{
		ValuePos: token.NoPos,
		Kind:     bl.Kind,
		Value:    bl.Value,
	}
}
