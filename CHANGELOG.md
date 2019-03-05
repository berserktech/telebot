# 0.1.0
* Rewritten in a modular manner.
  * `main`, the package at the root. It holds the Handle function that
    is used by Zeit. This function currenlty receives GitHub's
    Webhooks, parses the data, builds a formatted message and sends it
    through Telegram.
  * `gh` package for GitHub-related code.
  * `tg` package for Telegram-related code.
* Added missing go.mod files for telebot and it's submodules.
* Added a bash script to automate checking if the code is formatted:
  * `fmt-check.bash`: Runs `go fmt -l .` in all the submodules of this
    repo. If it finds files that don't have the proper formatting, it
    will exit with status code 1. We use this one for CI purposes
* Muted `synchronize` actions (specifically for PullRequestEvents).
* Made sure both CircleCI and Zeit worked fine with this new
  file structure.

# 0.0.1
* Initial version
