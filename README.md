# [ent](https://entgo.io/) - reproduce

> For querying, FilterProfileRule works on user.profiles, but the similar FilterTenantRule does NOT work on user.tenants

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
