# TODO

## Dependabot configuration improvements

Items identified via the `tune-dependabot-config` skill review on 2026-07-05.
Each item can be a standalone PR once bandwidth allows.

### Group tuning

- [ ] **Split gomod major bumps explicitly** — Add a `major-updates` group for `gomod` (mirroring the github-actions pattern) if Go modules start shipping coordinated major releases (e.g. `golang.org/x/...` family). Not needed today since the `x/net` module is the only production dep, but worth revisiting when the dependency count grows.
- [ ] **Per-family grouping for gomod** — If indirect deps start appearing, consider naming groups after import-path prefixes (`golang.org/x/*`, `github.com/stretchr/*`) so security advisories for one family don't block unrelated bumps.

### Cooldown tuning

- [ ] **Tighten patch cooldown for gomod** — The `gomod` ecosystem supports `semver-patch-days`. Once comfortable with grouping behaviour, consider `semver-patch-days: 3` and `semver-major-days: 14` alongside the current `default-days: 7` for finer-grained control. See [dependabot options reference](https://docs.github.com/en/code-security/dependabot/working-with-dependabot/dependabot-options-reference#cooldown--).
- [ ] **Security-update fast lane** — Add a `security-updates` entry per ecosystem with no cooldown if security response time SLO ever tightens.

### PR metadata

- [ ] **Add dependabot labels** — Add `labels: ["dependencies"]` (and optionally `["go"]` / `["github-actions"]`) to each ecosystem entry to make PR filtering easier in the GitHub UI.
- [ ] **Add assignees** — Add `assignees: ["jonasbn"]` so dependabot PRs land in the right review queue without relying on CODEOWNERS.

### Schedule

- [ ] **Stagger update schedules** — Both ecosystems currently run weekly with no `day:` constraint. Setting `day: "monday"` for github-actions and `day: "wednesday"` for gomod avoids a Monday pile-up if both trigger simultaneously.
- [ ] **Consider daily for gomod** — The module graph is small (2 direct deps); daily checks would catch CVEs faster. Evaluate after adding labels so the queue stays manageable.

### Open-pull-requests-limit

- [ ] **Set a PR cap** — Consider `open-pull-requests-limit: 5` per ecosystem to prevent accumulation during vacations or freeze windows. Remove the cap if automation is reliable enough to auto-merge.

### Security-only entries

- [ ] **Add dedicated security-updates entries** — Add separate `applies-to: security-updates` entries for both ecosystems with no cooldown and a tighter `open-pull-requests-limit`. This ensures CVE patches land fast even when version-updates are paused.
