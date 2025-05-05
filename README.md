THIS IS A MINI PROJECT FOR DEMO DOMAIN-DRIVEN DESIGN

STRUCTURE OF PROJECT

1. DOMAIN

Storing all the subdomains.
2. ENTITY

A struct that has an Identifier and that can change state.
3. VALUE OBJECTS

There can be occurrences where we have structs that are immutable and do not need a unique identifier.
Value objects are often found inside domains and used to describe certain aspects in that domain.
We will be creating one value object for now which is Transaction, once a transaction is performed, it cannot change state.
4. AGGREGATES

An Aggregate is a set of entities and value objects combined.
Reference:

https://dev.to/techschoolguru/how-to-setup-github-actions-for-go-postgres-to-run-automated-tests-81o
https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/
https://refactoring.guru/design-patterns/factory-method
