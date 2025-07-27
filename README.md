## Storage Query Tool for Gno.land

### Prerequisites

* `gnokey` is installed and in your `PATH`
* `gnoland` and `gnoweb` are running

### Overview
This program automates querying the storage and deposit usage for each realm
and package. It calculates the total and average values.

* Recursively fetches all realm and package paths from `/r/`  and `/p/`
* Runs `gnokey` queries for each
* Aggregates total and average `storage` and `deposit` across all realms

---
To run against local node

```bash
go run main.go
```

To check against remote network
```
go run main.go -web https://staging.gno.land -rpc https://rpc.staging.gno.land
```

---
To query storage deposit for individual contract uses gnokey instead:

```bash
gnokey query vm/qstorage --data "gno.land/r/demo/boards"
```

output

```
height: 0
data: storage: 89558, deposit: 8955800
```

