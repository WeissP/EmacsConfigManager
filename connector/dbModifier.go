package connector

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ResetTables() {
	conn := Connect()
	defer conn.Close()
	if _, err := conn.Exec(context.Background(), cmds["drop tables"]); err != nil {
		fmt.Println("Unable to drop tables")
		panic(err)
	}
	if _, err := conn.Exec(context.Background(), cmds["create tables"]); err != nil {
		fmt.Println("Unable to create tables")
		panic(err)
	}
}

func AddFileRecord(conn *pgxpool.Pool, module, class string) string {
	filename := fmt.Sprintf(filenameFormat, module, class)
	s := fmt.Sprintf(cmds["add file"], filename, module, class)
	if _, err := conn.Exec(context.Background(), s); err != nil {
		fmt.Println("Unable to add file")
		panic(err)
	}
	return filename
}

func AddParent(conn *pgxpool.Pool, child, parent string) {
	if IsChildExists(conn, child) {
		fmt.Printf("child %v has already parent:%s, new key is: %s", child, GetParent(conn, child), parent)
		return 
	}
	s := fmt.Sprintf(cmds["add parent"], child, parent)
	if _, err := conn.Exec(context.Background(), s); err != nil {
		fmt.Println("Unable to add parent")
		panic(err)
	}
}

func DeleteFile(conn *pgxpool.Pool, module, class string){
	s := fmt.Sprintf(cmds["delete file"], module, class)
	if _, err := conn.Exec(context.Background(), s); err != nil {
		fmt.Println("Unable to delete file")
		panic(err)
	}
}

func DeleteRelation(conn *pgxpool.Pool, child string){
	s := fmt.Sprintf(cmds["delete relation"], child)
	if _, err := conn.Exec(context.Background(), s); err != nil {
		fmt.Println("Unable to delete relation")
		panic(err)
	}
}


