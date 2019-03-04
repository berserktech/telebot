# 0.1.0
* Rewritten in a modular manner.
  * `gh` package for GitHub-related code.
  * `tg` package for Telegram-related code.
  * `handler` package for the actual handling function that's passed to Zeit.
* Added missing go.mod files for telebot and it's submodules.
* Added bash scripts for CI and automation processes:
  * `test.bash`: Runs `go test` in all the submodules of this repo.
  * `fmt.bash`: Runs `go fmt` in all the submodules of this repo.
  * `fmt-check.bash`: Runs `go fmt -l .` in all the submodules of this
    repo. If it finds files that don't have the proper formatting, it
    will exit with status code 1. We use this one for CI purposes
  * `mod-update.bash`: Updates all the versions of local dependencies of
    this module to the latest hash in your local git. Make sure to
    have pushed it. To reconcile with Go's authomatic changes, it runs
    `fmt.bash` at the end, which will also check if the version is
    available through the network. To use this, you must pass a
    git-reference to update to. You can use `HEAD`, as in the
    following example: `./mod-update.bash HEAD~1`.
* Muted `synchronized` actions (specifically for PullRequestEvents).

# 0.0.1
* Initial version
