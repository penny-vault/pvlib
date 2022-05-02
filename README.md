# pvlib

pvlib is a helper library for interfacing with the pv-api REST
API. It provides go types for serializing to/from colfer and methods
for uploading/downloading from pv-api.

# usage

```bash
go get github.com/penny-vault/pvlib
```

```go
pvlib.Authorize("<API Key>")
pvlib.FindPortfolio(pvlib.SearchCriteria{
    Name: "Name",
    Comparator: pvlib.EQUAL,
    StringValue: "My Portfolio"
})
```