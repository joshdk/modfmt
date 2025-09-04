module github.com/joshdk/modfmt

go 1.25.0

require (
	github.com/joshdk/buildversion v0.1.0
	github.com/joshdk/modfmt/pkg/modfmt v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.10.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	golang.org/x/mod v0.27.0 // indirect
)

replace (
	github.com/joshdk/modfmt/pkg/modfmt => ./pkg/modfmt
)
