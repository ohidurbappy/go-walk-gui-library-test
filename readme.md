### Install dependencies

```bash
go get github.com/lxn/walk
```

### Build

```bash
go build -ldflags="-H windowsgui"
```

### Manifest

```bash
go get github.com/akavel/rsrc
rsrc -manifest test.manifest -o rsrc.syso
```


or rename the test.manifest file to test.exe.manifest and distribute it with the application instead.