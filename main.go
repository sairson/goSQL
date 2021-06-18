package main

import (
	"SQLPrivilege/data"
	"SQLPrivilege/plugin"
	"flag"
	"fmt"
)
var (
	host string
	username string
	passowrd string
	port int
	module int
)

func init(){
	flag.StringVar(&host,"host","","hostname")
	flag.IntVar(&port,"port",1433,"port")
	flag.StringVar(&username,"username","sa","username")
	flag.StringVar(&passowrd,"password","123456","password")
	flag.IntVar(&module,"method",1,"Execute method" +
		"\n xp_cmdshell   1   <*>Echoable<*>" +
		"\n sp_oacreate   2   <*>Echoable<*>"+
		"\n open_clr      3   <*>No Echoable<*>")
	flag.Parse()
}

func main(){
	if host != ""{
		fmt.Println(" ██████╗  ██████╗ ███████╗ ██████╗ ██╗    ")
		fmt.Println("██╔════╝ ██╔═══██╗██╔════╝██╔═══██╗██║     ")
		fmt.Println("██║  ███╗██║   ██║███████╗██║   ██║██║ by SaiRson ")
		fmt.Println("██║   ██║██║   ██║╚════██║██║▄▄ ██║██║     ")
		fmt.Println("╚██████╔╝╚██████╔╝███████║╚██████╔╝███████╗")
		fmt.Println(" ╚═════╝  ╚═════╝ ╚══════╝ ╚══▀▀═╝ ╚══════╝")
		var OBJ data.MSSQLOBJ
		OBJ.Host = host
		OBJ.Port = port
		OBJ.Username = username
		OBJ.Password = passowrd
		fmt.Println(OBJ)
		switch module {
		case 1:
			fmt.Println("[+] start to run xp_cmdshell ")
			plugin.Mssql_xp_cmdshell(&OBJ)
			break
		case 2:
			fmt.Println("[+] start to run sp_oacreate ")
			plugin.Mssql_sp_oacreate(&OBJ)
			break
		case 3:
			fmt.Println("[+] start Try to Open CLR execute extend")
			plugin.OpenCLR(&OBJ)
		case 4:
			fmt.Println("[+] start to remove CLR execute extend")
			plugin.CloseCLR(&OBJ)
		default:
			fmt.Println(flag.ErrHelp)
		}
	}else{
		flag.Usage()
	}
}
