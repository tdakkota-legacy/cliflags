module github.com/tdakkota/cliflags/vksdk

go 1.15

require (
	github.com/SevereCloud/vksdk/v2 v2.9.0
	github.com/rs/zerolog v1.20.0
	github.com/tdakkota/cliflags v0.1.0
	github.com/tdakkota/vksdkutil/v2 v2.0.1
	github.com/urfave/cli/v2 v2.3.0
)

replace github.com/tdakkota/cliflags => ../.
