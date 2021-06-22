package main

import (
	"EmacsConfigManager/connector"
	"fmt"
	"os"
	"strings"
)

func main() {
	conn := connector.Connect()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			if len(os.Args) < 4 {
				printHelp()
				return 
			}
			switch os.Args[2] {
			case "file":
				fmt.Printf("%v", connector.AddFileRecord(conn,os.Args[3], os.Args[4]))
			case "parent":
				connector.AddParent(conn,os.Args[3], os.Args[4])
			default:
				printHelp()
				return 
			}
		case "exists":
			fmt.Printf("%v", connector.IsChildExists(conn, os.Args[2])) 
		case "reset":
			connector.ResetTables()
		case "delete":
			switch os.Args[2] {
			case "file":
				connector.DeleteFile(conn, os.Args[3], os.Args[4])
			case "relation":
				connector.DeleteRelation(conn, os.Args[3])
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
				for _, x := range connector.GetAllFilesByClass(conn,os.Args[3]){
					fmt.Printf("%v\n", x)
				}
			case "fuzzy":				
				for _, x := range connector.FuzzySearch(conn,os.Args[3], os.Args[4:]){
					fmt.Printf("%v\n", x)
				}
			case "all-parents":
				fmt.Printf("%v", strings.Join(connector.GetAllParents(conn, os.Args[3]), ","))
			case "child":
				fmt.Printf("%v", connector.GetChild(conn,os.Args[2]))
			default:
				printHelp()
				return 
			}
		default:
			
		}
	} else {
		fmt.Printf("%v", connector.GetAllParents(conn, "lsp-java"))
		fmt.Printf("%v", connector.SearchClasses(conn, "hl"))
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
