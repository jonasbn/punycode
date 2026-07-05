# Bundle Dependabot PRs + Fix Dependabot Config Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Bundle three split codeql-action dependabot PRs into one, merge two standalone bumps, harden the dependabot config with grouping + cooldown, and produce a TODO list of future improvements.

**Architecture:** Each task works on its own feature branch (direct commits to `main` are blocked by the `no-commit-to-branch` pre-commit hook). Tasks 1, 4, and 5 each open their own PR. Tasks 2 and 3 merge existing dependabot PRs via `gh pr merge`. Tasks can be executed in parallel where noted.

**Tech Stack:** Go CLI repo, GitHub Actions CI, GoReleaser, Dependabot, `gh` CLI, `python3` (YAML sanity check)

## Global Constraints

- Never commit directly to `main` — the `no-commit-to-branch` pre-commit hook will block it; always work on a named feature branch.
- All `gh` commands assume the `jonasbn/punycode` remote — do not add `--repo` unless the context is unclear.
- The existing `.github/dependabot.yml` uses double-quote string style throughout; new strings must use double quotes to match.
- Preserve all existing comments in `.github/dependabot.yml`.

---

### Task 1: Bundle codeql-action PRs #160, #162, #163 into a single PR

All three PRs update the same three steps in `.github/workflows/codeql.yml`, replacing the same old commit hash `8aad20d150bbac5944a9f9d289da16a4b0d87c1e` (v4.36.2) with `54f647b7e1bb85c95cddabcd46b0c578ec92bc1a` (v4.36.3).

**Files:**
- Modify: `.github/workflows/codeql.yml` (lines 47, 61, 74)

**Interfaces:**
- Produces: a new PR that supersedes and closes #160, #162, #163

- [ ] **Step 1: Create a feature branch**

```bash
git checkout main
git pull
git checkout -b chore/bundle-codeql-action-v4.36.3
```

- [ ] **Step 2: Apply all three hash replacements**

In `.github/workflows/codeql.yml`, replace every occurrence of the old hash with the new hash. All three lines change from `@8aad20d150bbac5944a9f9d289da16a4b0d87c1e # v4.36.2` to `@54f647b7e1bb85c95cddabcd46b0c578ec92bc1a # v4.36.3`:

Line 47 — init:
```yaml
      uses: github/codeql-action/init@54f647b7e1bb85c95cddabcd46b0c578ec92bc1a # v4.36.3
```

Line 61 — autobuild:
```yaml
      uses: github/codeql-action/autobuild@54f647b7e1bb85c95cddabcd46b0c578ec92bc1a # v4.36.3
```

Line 74 — analyze:
```yaml
      uses: github/codeql-action/analyze@54f647b7e1bb85c95cddabcd46b0c578ec92bc1a # v4.36.3
```

Run `sed` to do all three at once:
```bash
sed -i '' 's/8aad20d150bbac5944a9f9d289da16a4b0d87c1e # v4.36.2/54f647b7e1bb85c95cddabcd46b0c578ec92bc1a # v4.36.3/g' .github/workflows/codeql.yml
```

Verify — three lines should show `v4.36.3`, none should show `v4.36.2`:
```bash
grep -n "codeql-action" .github/workflows/codeql.yml
```
Expected:
```
47:      uses: github/codeql-action/init@54f647b7e1bb85c95cddabcd46b0c578ec92bc1a # v4.36.3
61:      uses: github/codeql-action/autobuild@54f647b7e1bb85c95cddabcd46b0c578ec92bc1a # v4.36.3
74:      uses: github/codeql-action/analyze@54f647b7e1bb85c95cddabcd46b0c578ec92bc1a # v4.36.3
```

- [ ] **Step 3: Commit**

```bash
git add .github/workflows/codeql.yml
git commit -m "chore: bump github/codeql-action from v4.36.2 to v4.36.3

Bundles dependabot PRs #160, #162, #163 — all three steps
(init, autobuild, analyze) use the same underlying action release."
```

- [ ] **Step 4: Push and open PR**

```bash
git push -u origin chore/bundle-codeql-action-v4.36.3
gh pr create \
  --title "chore: bump github/codeql-action from v4.36.2 to v4.36.3" \
  --body "$(cat <<'EOF'
Bundles dependabot PRs #160, #162, and #163 into a single PR.

All three steps (`init`, `autobuild`, `analyze`) are part of the same `github/codeql-action` release and should always move together.

Closes #160
Closes #162
Closes #163
EOF
)"
```

- [ ] **Step 5: Close the three superseded dependabot PRs**

After the new PR is open:
```bash
gh pr close 160 --comment "Superseded by the bundled PR that upgrades all three codeql-action steps together."
gh pr close 162 --comment "Superseded by the bundled PR that upgrades all three codeql-action steps together."
gh pr close 163 --comment "Superseded by the bundled PR that upgrades all three codeql-action steps together."
```

---

### Task 2: Merge PR #161 — spellcheck-github-actions bump (standalone)

PR #161 bumps `rojopolis/spellcheck-github-actions` from 0.62.0 to 0.63.0. This is an independent tool; no bundling needed.

**Files:** None to edit locally — this is a direct merge of the existing dependabot PR.

- [ ] **Step 1: Review the PR diff**

```bash
gh pr diff 161
```

Confirm the change touches only `.github/workflows/spellcheck.yml` and updates only the version/hash for `rojopolis/spellcheck-github-actions`.

- [ ] **Step 2: Merge the PR**

```bash
gh pr merge 161 --squash --delete-branch
```

- [ ] **Step 3: Verify**

```bash
gh pr view 161 --json state --jq '.state'
```
Expected: `MERGED`

---

### Task 3: Merge PR #164 — goreleaser-action bump (standalone)

PR #164 bumps `goreleaser/goreleaser-action` from 7.2.2 to 7.2.3. This is an independent tool; no bundling needed.

**Files:** None to edit locally — this is a direct merge of the existing dependabot PR.

- [ ] **Step 1: Review the PR diff**

```bash
gh pr diff 164
```

Confirm the change touches only `.github/workflows/release.yml` and updates only the version/hash for `goreleaser/goreleaser-action`.

- [ ] **Step 2: Merge the PR**

```bash
gh pr merge 164 --squash --delete-branch
```

- [ ] **Step 3: Verify**

```bash
gh pr view 164 --json state --jq '.state'
```
Expected: `MERGED`

---

### Task 4: Harden dependabot config with groups and cooldown

Fix `.github/dependabot.yml` so future codeql-action bumps (and all other GitHub Actions updates) are automatically grouped, and a 7-day cooldown prevents broken releases from landing immediately.

**Files:**
- Modify: `.github/dependabot.yml`

**Interfaces:**
- Produces: a PR with the updated config; once merged, dependabot will batch same-ecosystem updates automatically

- [ ] **Step 1: Create a feature branch**

```bash
git checkout main
git pull
git checkout -b chore/harden-dependabot-config
```

- [ ] **Step 2: Rewrite `.github/dependabot.yml`**

Replace the entire file with the following (preserving all existing comments):

```yaml
version: 2
updates:
  - package-ecosystem: "github-actions"
    # Look for `.github/workflows` in the `root` directory
    directory: "/"
    # Check for updates once a week
    schedule:
      interval: "weekly"
    groups:
      major-updates:
        patterns:
          - "*"
        update-types:
          - "major"
      minor-and-patch:
        patterns:
          - "*"
        update-types:
          - "minor"
          - "patch"
    cooldown:
      default-days: 7

  - package-ecosystem: "gomod"
    # Look for Go code in the `root` directory
    directory: "/"
    # Check for updates once a week
    schedule:
      interval: "weekly"
    groups:
      minor-and-patch:
        patterns:
          - "*"
        update-types:
          - "minor"
          - "patch"
    cooldown:
      default-days: 7
```

Why two groups for `github-actions`: `init`, `autobuild`, and `analyze` ship coordinated releases — merging one without the others breaks the CodeQL workflow. Batching major bumps together keeps the coordinated set atomic. Minor/patch updates get their own group so they can land without waiting for major-bump review.

Why one group for `gomod`: Go module major versions are separate import paths and breaking changes that need individual review; minor/patch bumps are safe to batch.

- [ ] **Step 3: Run the YAML sanity check**

```bash
python3 -c 'import yaml,sys; yaml.safe_load(open(".github/dependabot.yml")); print("OK")'
```
Expected: `OK`

- [ ] **Step 4: Commit**

```bash
git add .github/dependabot.yml
git commit -m "chore: add dependabot groups and 7-day cooldown

Groups minor+patch github-actions updates into a single rolling PR,
adds a separate group for coordinated major-action bumps, and batches
gomod minor+patch updates. The 7-day cooldown prevents yanked releases
from landing before maintainers retract them."
```

- [ ] **Step 5: Push and open PR**

```bash
git push -u origin chore/harden-dependabot-config
gh pr create \
  --title "chore: add dependabot groups and 7-day cooldown" \
  --body "$(cat <<'EOF'
## Summary

- Adds `major-updates` and `minor-and-patch` groups to `github-actions` ecosystem so related action steps (init/autobuild/analyze) are always bumped in a single PR rather than one PR per step
- Adds `minor-and-patch` group to `gomod` ecosystem so routine Go dependency bumps are batched
- Adds `cooldown: default-days: 7` to both ecosystems to let the community retract broken releases before dependabot files a PR

## Why two groups for github-actions?

CodeQL's `init`, `autobuild`, and `analyze` steps ship coordinated releases — the artifact format can change across steps. Major bumps need to land together or CI breaks. Minor/patch updates are safe to batch separately so they don't wait on major-bump review.

## Test plan
- [ ] Confirm `.github/dependabot.yml` parses: `python3 -c 'import yaml; yaml.safe_load(open(".github/dependabot.yml")); print("OK")'`
- [ ] After merge, watch the next dependabot run — expect grouped PRs instead of per-action individual PRs
EOF
)"
```

---

### Task 5: Create docs/TODO.md with future dependabot improvements

Generate a `docs/TODO.md` file capturing the backlog of dependabot configuration improvements that are out of scope for today but worth tracking.

**Files:**
- Create: `docs/TODO.md`

**Interfaces:**
- Produces: a committed `docs/TODO.md` on a feature branch, opened as a PR

- [ ] **Step 1: Create the docs directory and TODO file**

Create `docs/TODO.md` with the following content:

```markdown
# TODO

## Dependabot configuration improvements

Items identified via the `tune-dependabot-config` skill review on 2026-07-05.
Each item can be a standalone PR once bandwidth allows.

### Group tuning

- [ ] **Split gomod major bumps explicitly** — Add a `major-updates` group for `gomod` (mirroring the github-actions pattern) if Go modules start shipping coordinated major releases (e.g. `golang.org/x/...` family). Not needed today since the `x/net` module is the only dep, but worth revisiting when the dependency count grows.
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
```

- [ ] **Step 2: Create a feature branch and commit**

```bash
git checkout main
git pull
git checkout -b docs/dependabot-todo
git add docs/TODO.md
git commit -m "docs: add dependabot TODO list from tune-dependabot-config review"
```

- [ ] **Step 3: Push and open PR**

```bash
git push -u origin docs/dependabot-todo
gh pr create \
  --title "docs: add dependabot TODO list from config review" \
  --body "$(cat <<'EOF'
Adds `docs/TODO.md` with a backlog of dependabot configuration improvements
identified during the 2026-07-05 tune-dependabot-config skill review.

Items are scoped to: group tuning, cooldown tuning, PR metadata, schedule
adjustments, open-PR limits, and security-only entries. Each item is a
self-contained future PR.
EOF
)"
```

---

## Execution Order

Tasks 2 and 3 (standalone PR merges) can run in parallel with each other and with Task 1. Task 4 and Task 5 are independent of all other tasks and can also run in parallel. Recommended order to avoid context-switching overhead:

1. **Tasks 2 + 3** — merge the two standalone PRs (fast, no local changes needed)
2. **Task 1** — bundle and open the codeql PR, then close the three superseded PRs
3. **Tasks 4 + 5** — open the dependabot config hardening PR and the TODO doc PR

---

## Self-Review

**Spec coverage:**
- Bundle #160, #162, #163 → Task 1 ✓
- Merge #161 standalone → Task 2 ✓
- Merge #164 standalone → Task 3 ✓
- Fix dependabot config with groups → Task 4 ✓
- tune-dependabot-config TODO list in docs/TODO.md → Task 5 ✓

**Placeholder scan:** No TBDs, no "implement later", no steps without concrete commands or content.

**Type consistency:** No shared function signatures across tasks; each task is self-contained shell/YAML work.
