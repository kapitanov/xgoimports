LOCAL = local.example.com

TESTCASES = simple \
	already-sorted \
	nothing-to-sort \
	few-import-groups \
	misformatted-by-ide \
	dot-import \
	discarded-import \
	aliased-import \
	import-spec-comments \
	import-group-comments \
	cgo-import \
	issue-1

all:
	@$(foreach testcase,$(TESTCASES),make _test TESTCASE=$(testcase);)

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
