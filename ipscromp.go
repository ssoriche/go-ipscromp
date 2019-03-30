package main // import "github.com/ssoriche/go-ipscromp"

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"github.com/spf13/viper"
	"net"
	"strings"
)

func main() {
	viper.SetConfigName("ipscromp")       // name of config file (without extension)
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	conn, err := net.Dial("tcp", viper.GetString("server"))
	if err != nil {

	}

	fmt.Fprint(conn, "USER "+viper.GetString("user")+" 2")
	recv, err := bufio.NewReader(conn).ReadString('\n')
	token := strings.TrimSpace(recv[4:])

	sum := sha1.Sum([]byte(viper.GetString("user") + ":" + token + ":" + viper.GetString("password")))
	fmt.Fprintf(conn, "PERMIT %x", sum)
	recv, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Firewall Opened")
	}
	fmt.Println(recv)
}
