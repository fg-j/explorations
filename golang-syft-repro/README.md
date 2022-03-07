# Missing Go Module `artifacts` when modules are vendored

Reproduction steps:
1. Run `./repro.sh`
2. See that the resulting `syft` output has no `artifacts`
3. Comment out the lines that vendor the dependencies
4. Re-run `./repro.sh`
5. See that the resulting `syft` output has `artifacts`
