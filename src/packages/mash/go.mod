module mash.com/mash

require (
	mash.com/event v0.0.0
	mash.com/transport v0.0.0
)

replace (
	mash.com/event => ../event
	mash.com/transport => ../transport
)

go 1.12
