package plugin

import (
	"SQLPrivilege/data"
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"os"
)

func ConnDB(mssqlobj *data.MSSQLOBJ)(*sql.DB, error){
	ConnString := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v",
		mssqlobj.Host,
		mssqlobj.Port,
		mssqlobj.Username,
		mssqlobj.Password,
		"master",
	)
	conn, err := sql.Open("mssql",ConnString)
	if err != nil {
		return nil, err
	}
	return conn,nil
}

func CMD_SHELL() string{
	fmt.Printf("[+] shell >> ")
	var str string
	//使用os.Stdin开启输入流
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		str = in.Text()
	}
	return str
}

func SQLQuery(query string,conn *sql.DB) (*sql.Stmt,error){
	Query, err := conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	return Query ,nil
}


