package connector

import (
	"database/sql"
	"fmt"
	"strings"
)

func GetParent(db *sql.DB, child string) (parent string) {
	s := fmt.Sprintf(cmds["get parent"], child)
	err := db.QueryRow(s).Scan(&parent)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
		panic(err)
	}
	return parent
}

func IsParentExists(db *sql.DB, parent string) bool {
	s := fmt.Sprintf(cmds["is parent exists"], parent)
	var r string
	err := db.QueryRow(s).Scan(&r)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		panic(err)
	}
	return true
}

func GetAllParents(db *sql.DB, child string) (parents []string) {
	parent := GetParent(db, child)
	if parent == "" {
		return parents
	}
	return append(GetAllParents(db, parent), parent)
}

func GetChild(db *sql.DB, parent string) (child string) {
	s := fmt.Sprintf(cmds["get child"], child)
	err := db.QueryRow(s).Scan(&child)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
		panic(err)
	}
	return child
}

func IsChildExists(db *sql.DB, child string) bool {
	s := fmt.Sprintf(cmds["is child exists"], child)
	// fmt.Printf("%v", s)
	var r string
	err := db.QueryRow(s).Scan(&r)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		panic(err)
	}
	return true
}

func GetAllChildren(db *sql.DB, parent string) (children []string) {
	child := GetChild(db, parent)
	if child == "" {
		return children
	}
	return append(GetAllChildren(db, child), child)
}

func searchChildren(db *sql.DB, child string) (children []string) {
	q := fmt.Sprintf(cmds["search children"], child)
	rows, _ := db.Query(q)
	for rows.Next() {
		var (
			c string
		)
		err := rows.Scan(&c)
		if err != nil {
			panic(err)
		}
		children = append(children, c)
	}
	return children
}

func searchParents(db *sql.DB, parent string) (parents []string) {
	q := fmt.Sprintf(cmds["search parents"], parent)
	rows, _ := db.Query(q)
	for rows.Next() {
		var (
			p string
		)
		err := rows.Scan(&p)
		if err != nil {
			panic(err)
		}
		parents = append(parents, p)
	}
	return parents
}

func SearchClasses(db *sql.DB, class string) (classes []string) {
	q := fmt.Sprintf(cmds["search classes"], class)
	rows, _ := db.Query(q)
	for rows.Next() {
		var class string
		err := rows.Scan(&class)
		if err != nil {
			panic(err)
		}
		classes = append(classes, class)
	}
	return classes
}

func SearchParentsAndChildren(db *sql.DB, class string) (classes []string) {
	qp := fmt.Sprintf(cmds["search parents"], class)
	qc := fmt.Sprintf(cmds["search children"], class)
	q := fmt.Sprintf("select * from ((%v) union (%v)) as mixed", qp, qc)
	rows, _ := db.Query(q)
	for rows.Next() {
		var c string
		err := rows.Scan(&c)
		if err != nil {
			panic(err)
		}
		classes = append(classes, c)
	}
	return classes
}

func FuzzySearch(db *sql.DB, module string, classes []string) (files []string) {
	// classes has mutiple c's, use c to get its similiar names in class_relation (as cc's), each cc has multiple children ccc
	condArray := []string{}
	for _, c := range classes {
		cc := SearchClasses(db, c)
		// fmt.Printf("cc:%v\n", cc)
		for j, x := range cc {
			ccc := append(GetAllChildren(db, x), x)
			for k, y := range ccc {
				ccc[k] = fmt.Sprintf("file.class like '%%%v%%'", y)
			}
			cc[j] = strings.Join(ccc, " OR ")
		}
		condArray = append(condArray, strings.Join(cc, " OR "))
	}
	var cond string
	if len(condArray) == 0 {
		cond = ""
	} else {
		cond = fmt.Sprintf("AND (%v)", strings.Join(condArray, " AND "))
	}
	q := fmt.Sprintf("select filename from file where (module_name like '%%%s%%') %v\n", module, cond)
	// fmt.Printf("%v", q)
	rows, _ := db.Query(q)
	for rows.Next() {
		var file string
		err := rows.Scan(&file)
		if err != nil {
			panic(err)
		}
		files = append(files, file)
	}
	return files
}

func GetAllFilesByClass(db *sql.DB, class string) (files []string) {
	q := fmt.Sprintf(cmds["search files by class"], class)
	rows, _ := db.Query(q)
	for rows.Next() {
		var file string
		err := rows.Scan(&file)
		if err != nil {
			panic(err)
		}
		files = append(files, file)
	}
	return files
}
