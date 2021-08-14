all: dnazart

dnazart:
	@./scripts/build.sh dnazart

package:
	@rm -rf packages/
	@./scripts/package.sh Linux   linux   amd64
	@./scripts/package.sh Mac     darwin  amd64
	@./scripts/package.sh Windows windows amd64