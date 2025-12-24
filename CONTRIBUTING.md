# Contributing

## Reporting Issues

Before submitting a new issue, please search for an existing or similar issue.

## Pull Requests

Pull requests are always welcome and very much appreciated. However, there are some lightweight
guidelines below to make reviewing easier, faster, smoother, and more consistent.

If a pull request has been waiting too long for a review, please tag the reviewer to draw attention
to it.

### Titles

Pull request titles must conform to the [Conventional Commits][1] specification. This means the
title should be in the format `{type}: {subject}`. The subject should be a short summary of the
change, whilst type should be one of the following:

| Type       | Description                                                   |
| ---------- | ------------------------------------------------------------- |
| `chore`    | Other changes that don't modify code files.                   |
| `ci`       | Changes to the CI configuration files.                        |
| `docs`     | Documentation only changes.                                   |
| `feat`     | A new feature.                                                |
| `fix`      | Bug fixes.                                                    |
| `perf`     | A performance improvement.                                    |
| `refactor` | A code change that is neither a bug fix nor a new feature.    |
| `revert`   | Reverts a previous commit.                                    |
| `style`    | Formatting changes that do no affect the meaning of the code. |
| `test`     | Adding missing tests or fixing existing tests.                |

All commits in the pull request are squashed on merge, with the PR title and description being used
for the `main` branch commit. Pull request titles will also be used to create release notes.

[1]: https://www.conventionalcommits.org/

### Descriptions

Pull request descriptions should fill out the provided template, replacing the placeholder text.

Please ensure the rationale of the change is included, including the decision making process behind
it. This can be particularly useful if the decision needs to be revisited in the weeks or months
after the change was originally made.

### Size

Ideally, a pull request will focus on a single change. This helps reviewers avoid context switching
in a single review. If the pull request description includes the word "also", it may be worth
splitting the change into 2 or more pull requests. If you are unsure, please ask the reviewers
before splitting.

Pull requests are labelled according to size. Ideally a change will be categorised as small or
medium. If the pull request is categorised as large, extra-large or extra-extra-large, reviewers may
ask that it be split up. However, larger pull requests may still be accepted if there is no good
place to split the change.
