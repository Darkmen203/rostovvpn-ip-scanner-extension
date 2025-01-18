package main

import (
	"github.com/Darkmen203/rostovvpn-core/extension/server"
	_ "github.com/Darkmen203/rostovvpn-ip-scanner-extension/hiddify_extension"
)

func main() {
	server.StartTestExtensionServer()

}
