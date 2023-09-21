## Steps to reproduce:

<from project root>: go run test/test.go

## Notes:
The schema has two types, Test and Test2. Test2 is an optional single link of Test. When the golang code is generated, it does not embed `edgedb.Optional` into the struct for Test2. This means any queries for a Test object returns the following error;

```
edgedb.InvalidArgumentError: the "out" argument does not match query schema: expected queries.testQueryResulttest_2Item at queries.testQueryResult.test_2 to be edgedb.Optional because the field is not required
```

I have attached a modified version of the same generated code that works.
