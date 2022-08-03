# [ent](https://entgo.io/) - reproduce

This repo connectes to this [issue](https://github.com/ent/ent/issues/2827)

## The issue

> In short, it's the FilterFunc works for user.profiles but NOT triggering at all for user.tenants.

## Structure

- ent:<br/>
  --schema: all ent schemas<br/>
  --rules: all ent privacy rules

- entgen: ent generated code
- tests: golang tests for reproducing the issue

## How to reproduce

run tests: `go test ./tests/... -v`

- `TestUserShouldNotBeAbleToGetNonSameTenantProfiles()` is passing
- `TestUserShouldNotBeAbleToGetNonSameTenantTenants()` is NOT passing
