module mash.com/cli

go 1.12

require (
	github.com/urfave/cli v1.21.0
	mash.com/event v0.0.0
	mash.com/mash v0.0.0
	mash.com/transport v0.0.0
)

replace (
	mash.com/event => ./../packages/event
	mash.com/mash => ./../packages/mash
	mash.com/transport => ./../packages/transport
)
