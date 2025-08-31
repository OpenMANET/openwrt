/*
Package uci implements a binding to OpenWrt's UCI (Unified Configuration
Interface) files in pure Go.

The typical use case is reading and modifying UCI config options:
	import "github.com/digineo/go-uci"

	uci.Get("network", "lan", "ifname") //=> []string{"eth0.1"}, true
	uci.Set("network", "lan", "ipaddr", "192.168.7.1")
	uci.Commit() // or uci.Revert()

For more details head over to the OpenWrt wiki, or dive into UCI's C
source code:
 - https://openwrt.org/docs/guide-user/base-system/uci
 - https://git.openwrt.org/?p=project/uci.git;a=summary

The lexer is heavily inspired by Rob Pike's 2011 GTUG Sydney talk
"Lexical Scanning in Go" (https://talks.golang.org/2011/lex.slide,
https://youtu.be/HxaD_trXwRE), which in turn was a presentation of
an early version of Go's text/template parser. It follows, that this
library borrows code from Go's standard library (BSD-style licensed).

The UCI grammar (for the purpose of this library) is defined as follows:

	uci
			packageDecl*
			configDecl*

	packageDecl
			`package` value CRLF configDecl*

	configDecl
			`config` ident value? CRLF optionDecl*

	optionDecl
			`option` ident value
			`list` ident value

	ident
			[_a-zA-Z0-9]+

	value
			`'` STRING `'`
			`"` STRING `"`
			ident

For now, UCI imports/exports (packageDecl production) are not supported
yet. The STRING token (value production) is also somewhat vaguely
defined, and needs to be aligned with the actual C implementation.
*/
package uci
