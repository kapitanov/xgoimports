LOCAL = local.example.com/group/project

all:
	@make _test TESTCASE=simple
	@make _test TESTCASE=already-sorted
	@make _test TESTCASE=nothing-to-sort
	@make _test TESTCASE=few-import-groups
	@make _test TESTCASE=misformatted-by-ide
	@make _test TESTCASE=dot-import
	@make _test TESTCASE=discarded-import
	@make _test TESTCASE=aliased-import
	@make _test TESTCASE=import-spec-comments
	@make _test TESTCASE=import-group-comments
	@make _test TESTCASE=cgo-import

_test:
	@../artifacts/xgoimports -local $(LOCAL) -e ./testdata/$(TESTCASE)/input.txt > ./testdata/$(TESTCASE)/actual.txt
	@diff ./testdata/$(TESTCASE)/expected.txt ./testdata/$(TESTCASE)/actual.txt > ./testdata/$(TESTCASE)/expected.diff
	@DIFF=$(cat ./testdata/$(TESTCASE)/expected.diff); \
	if [ -n "$$DIFF" ]; then \
		printf "%bFAIL %s%b\n" "\033[31m" "$(TESTCASE)" "\033[0m" >&2; \
		exit 1; \
	else \
		printf "%bPASS %s%b\n" "\033[32m" "$(TESTCASE)" "\033[0m" >&2; \
	fi
