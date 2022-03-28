module github.com/yoeria/iwdConversion-go

go 1.18

require (
	github.com/lukegb/dds v0.0.0-20190402175749-8b7170e64003
	github.com/spf13/cobra v1.4.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace (
	github.com/yoeria/iwdConversion-go/cmd => ./cmd
	github.com/yoeria/iwdConversion-go/hash => ./hash
)
