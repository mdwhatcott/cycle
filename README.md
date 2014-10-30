There's a bug in the cover tool. Run the following command in this package:

    go test -cover

That command never exits (at least I've never waited long enough to see it exit). It looks like having an import cycle in your test files causes the go cover tool to hang. Now, I know, you shouldn't run coverage on a package that has an import cycle. Unfortunately, the only way to discover that you have an import cycle is to run go test. Any other build failure is reported when invoking the `go test -cover`, so why not an import cycle?