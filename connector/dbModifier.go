package connector

import (
	"database/sql"
	"fmt"
)

func ResetTables() {
	db := Connect()
	defer db.Close()
	if _, err := db.Exec(cmds["drop tables"]); err != nil {
		fmt.Println("Unable to drop tables")
		panic(err)
	}
	if _, err := db.Exec(cmds["create tables"]); err != nil {
		fmt.Println("Unable to create tables")
		panic(err)
	}
}

func AddFileRecord(db *sql.DB, module, class string) string {
	filename := fmt.Sprintf(filenameFormat, module, class)
	s := fmt.Sprintf(cmds["add file"], filename, module, class)
	if _, err := db.Exec(s); err != nil {
		fmt.Println("Unable to add file")
		panic(err)
	}
	return filename
}

func AddParent(db *sql.DB, child, parent string) {
	if IsChildExists(db, child) {
		fmt.Printf("child %v has already parent:%s, new key is: %s", child, GetParent(db, child), parent)
		return 
	}
	s := fmt.Sprintf(cmds["add parent"], child, parent)
	if _, err := db.Exec(s); err != nil {
		fmt.Println("Unable to add parent")
		panic(err)
	}
}

func DeleteFile(db *sql.DB, module, class string){
	s := fmt.Sprintf(cmds["delete file"], module, class)
	if _, err := db.Exec(s); err != nil {
		fmt.Println("Unable to delete file")
		panic(err)
	}
}

func DeleteRelation(db *sql.DB, child string){
	s := fmt.Sprintf(cmds["delete relation"], child)
	if _, err := db.Exec(s); err != nil {
		fmt.Println("Unable to delete relation")
		panic(err)
	}
}


