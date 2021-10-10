package main

import (
	"EmacsConfigManager/connector"
	"fmt"
	"os"
	"strings"
)

func main() {
	db := connector.Connect()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			if len(os.Args) < 4 {
				printHelp()
				return 
			}
			switch os.Args[2] {
			case "file":
				fmt.Printf("%v", connector.AddFileRecord(db,os.Args[3], os.Args[4]))
			case "parent":
				connector.AddParent(db,os.Args[3], os.Args[4])
			default:
				printHelp()
				return 
			}
		case "exists":
			fmt.Printf("%v", connector.IsChildExists(db, os.Args[2])) 
		case "reset":
			connector.ResetTables()
		case "delete":
			switch os.Args[2] {
			case "file":
				connector.DeleteFile(db, os.Args[3], os.Args[4])
			case "relation":
				connector.DeleteRelation(db, os.Args[3])
			default:
				printHelp()
				return 
			}			
		case "get":
			if len(os.Args) < 3 {
				return
			}
			switch os.Args[2] {
			case "files":				
				for _, x := range connector.GetAllFilesByClass(db,os.Args[3]){
					fmt.Printf("%v\n", x)
				}
			case "fuzzy":				
				for _, x := range connector.FuzzySearch(db,os.Args[3], os.Args[4:]){
					fmt.Printf("%v\n", x)
				}
			case "all-parents":
				fmt.Printf("%v", strings.Join(connector.GetAllParents(db, os.Args[3]), ","))
			case "child":
				fmt.Printf("%v", connector.GetChild(db,os.Args[2]))
			default:
				printHelp()
				return 
			}
		default:
			
		}
	} else {
		fmt.Printf("%v", connector.GetAllParents(db, "lsp-java"))
		fmt.Printf("%v", connector.SearchClasses(db, "hl"))
		// connector.AddFileRecord(conn, "settings", "hl-line")
		// connector.AddParent(conn, "hl-line", "ui")
		// connector.AddParent(conn, "popwin", "window-frame")
		// connector.AddParent(conn, "window-frame", "ui")
	}
	
}

func printHelp() {
	fmt.Print(`EmacsConfigManager
Usage:
  EmacsConfigManager add
  todo add task
  todo update task_num item
  todo remove task_num
Example:
  todo add 'Learn Go'
  todo list
`)
}
