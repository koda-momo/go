module trueserver

go 1.20

replace trueserver/functions => ./functions

require (
	trueserver/data v0.0.0-00010101000000-000000000000
	trueserver/functions v0.0.0-00010101000000-000000000000
)

replace trueserver/data => ./data
