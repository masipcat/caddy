package standard

import (
	// standard Caddy HTTP app modules
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/caddyauth"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/encode"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/encode/gzip"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/encode/zstd"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/fileserver"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/headers"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/map"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/push"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/requestbody"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/reverseproxy"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/reverseproxy/fastcgi"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/rewrite"
	_ "github.com/masipcat/caddy/v2/modules/caddyhttp/templates"
)
