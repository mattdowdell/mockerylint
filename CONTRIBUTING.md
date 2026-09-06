# Contributing

## Reporting Issues

Before submitting a new issue, please search for an existing or similar issue.

## Pull Requests

Pull requests are always welcome and very much appreciated. However, there are some lightweight
guidelines below to make reviewing easier, faster, smoother, and more consistent.

If a pull request has been waiting too long for a review, please tag the reviewer to draw attention
to it.

Each pull request is a collaborative process intended to improve the codebase and understanding of
both authors and reviewers. Questions can and should be asked where clarity is needed. The answers
to these questions may be worth adding to the code to aid future contributors.

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

If the template was accidentally deleted, it can be found [here][2].

[2]: ./.github/PULL_REQUEST_TEMPLATE.md

### Size

Ideally, a pull request will focus on a single change. This helps reviewers avoid context switching
in a single review. If the pull request description includes the word "also", it may be worth
splitting the change into 2 or more pull requests. If you are unsure, please ask the reviewers
before splitting.

Pull requests are labelled according to size. Ideally a change will be categorised as small or
medium. If the pull request is categorised as large, extra-large or extra-extra-large, reviewers may
ask that it be split up. However, larger pull requests may still be accepted if there is no good
place to split the change.

## Use of Generative AI

This project does not restrict the use of generative AI during development, maintenance or
documentation. However, regardless of the tools used, contributors remain responsible for everything
they submit. In short, abdication of responsibility to generative AI is both unacceptable and
unwelcome.

Issues and pull requests remain a human-centric process. All contributors are expected to make
their own choices and form their own opinions. Furthermore, each contributor can reasonably expect
that they are collaborating and communicating with another human.

The target audience for all documentation in the project is other humans. This includes code
comments, Markdown files, issues, pull request titles and pull request descriptions. Documentation
must be concise, well-structured and informative with that audience in mind.

For pull requests in particular, authors are expected to be able to explain their changes in their
own words. If an author cannot explain why the change is correct, it is not ready for review.
To reduce the burden on reviewers, a pull request description must not contain a verbose enumeration
of the diff or sections beyond what the template prescribes.

Finally, contributions must be yours to license under this project's terms. There is a vast body of
prior art for writing analyzers and much of it is worth learning from, but not all of it is
licensed compatibly. Generative AI will readily reproduce code it has seen, so you must check
that what you submit is yours to give.
