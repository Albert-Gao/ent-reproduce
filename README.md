# [ent](https://entgo.io/) - reproduce

> For querying, FilterProfileRule works on user.profiles, but the similar FilterTenantRule does NOT work on user.tenants

## Structure

- ent/schema : all ent schemas
- ent/rules : all ent privacy rules

## How to reproduce

run tests: `go test ./tests/... -v`

- `TestUserShouldNotBeAbleToGetNonSameTenantProfiles()` is passing
- `TestUserShouldNotBeAbleToGetNonSameTenantTenants()` is NOT passing
