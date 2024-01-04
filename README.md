# Experimenting

# Workflows

## adding a dependency

1. update go code: to include the new dependency

```go
# file:example/hello/reverse_test.go

# import the new dependency
import "gotest.tools/v3/assert"

...

# use the new dependency
assert.Equal(t, c.want, got)
````

2. run `go mod tidy` in the example/hello directory. This updates go.mod / go.sum
3. run gazelle

# What works

- [x] at "head" development for local dependencies (hello depends on example/hello for reverse.String and the local reverse.Int)
    - [x] VSCode Debugger
    - [x] VSCode jump to definition
    - [x] VSCode inline Tests
- [x] external dependency (github.com/stretchr/testify/assert)

# Current Known Gaps
- [ ] multiple seperate go.mod files with seperate dependencies
    - [ ] There are several ways to implement this, mostly it is important that we ensure bazel + native go are kept in sync. We want to avoid a scenario were a module is published in an unusable state due to it depending on a dependency it fails to specify, but rather gets via a sibling module which does.
 
# interesting and maybe relavant links

* https://bazel.build/external/overview
* https://github.com/bazelbuild/bazel/issues/19301
* https://github.com/bazelbuild/bazel/issues/20364
* https://docs.google.com/document/d/1dj8SN5L6nwhNOufNqjBhYkk5f-BJI_FPYWKxlB3GAmA/edit#heading=h.5mcn15i0e1ch
* https://github.com/bazelbuild/bazel/issues/17493
