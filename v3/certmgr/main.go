// certmgr is a tool for monitoring and renewing TLS certificates, and
// restarting services that need to use this.
package main

import "github.com/cloudflare/certmgr/v3/certmgr/cmd"

func main() {
	cmd.Execute()
}
