# Contributing to Shared-Sl

Thanks for your interest! OpenCroupier is maintained by **Chefmine**
(<ChMi8@pm.me>). Contributions are welcome through issues and pull/merge requests.

> The project is in an early design phase: opening an issue to discuss an idea
> before writing code is strongly encouraged.

## Reporting bugs and requesting features

Open an issue and include:

- what you expected and what happened,
- steps to reproduce,
- versions and hardware.

Do **not** publish security vulnerabilities in public issues. Email the maintainer instead.

## Development workflow

1. Fork the repository (or create a branch if you have access).
2. Create a branch from `main`: `feat/short-description` or `fix/short-description`.
3. Make focused changes; one logical change per pull request.
4. Add or update tests and documentation when relevant.
5. Open a pull/merge request against `main` and describe what and why.

`main` is protected: changes land only through reviewed pull/merge requests.

## Commit messages

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(server): add table assignment endpoint
fix(firmware): ignore repeated card reads for 3 seconds
docs: clarify pairing flow
```

Common types: `feat`, `fix`, `docs`, `refactor`, `test`, `chore`.
Common scopes: `server`, `firmware`, `ui`.

## Code style

- **All components**: no formatting convention is enforced. The only rule is that every function has one precise purpose; do not write functions that do several unrelated things.

## Security rules

- Keep all terminal ↔ server traffic over HTTPS.
