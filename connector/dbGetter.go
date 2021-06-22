package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetParent(conn *pgxpool.Pool, child string) (parent string) {
	s := fmt.Sprintf(cmds["get parent"], child)
	err := conn.QueryRow(context.Background(), s).Scan(&parent)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ""
		}
		panic(err)
	}
	return parent
}

func IsParentExists(conn *pgxpool.Pool, parent string) bool {
	s := fmt.Sprintf(cmds["is parent exists"], parent)
	var r string
	err := conn.QueryRow(context.Background(), s).Scan(&r)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		}
		panic(err)
	}
	return true
}

func GetAllParents(conn *pgxpool.Pool, child string) (parents []string) {
	parent := GetParent(conn, child)
	if parent == "" {
		return parents
	}
	return append(GetAllParents(conn, parent), parent)
}

func GetChild(conn *pgxpool.Pool, parent string) (child string) {
	s := fmt.Sprintf(cmds["get child"], child)
	err := conn.QueryRow(context.Background(), s).Scan(&child)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ""
		}
		panic(err)
	}
	return child
}

func IsChildExists(conn *pgxpool.Pool, child string) bool {
	s := fmt.Sprintf(cmds["is child exists"], child)
	// fmt.Printf("%v", s)
	var r string
	err := conn.QueryRow(context.Background(), s).Scan(&r)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		}
		panic(err)
	}
	return true
}

func GetAllChildren(conn *pgxpool.Pool, parent string) (children []string) {
	child := GetChild(conn, parent)
	if child == "" {
		return children
	}
	return append(GetAllChildren(conn, child), child)
}

func searchChildren(conn *pgxpool.Pool, child string) (children []string) {
	q :=  fmt.Sprintf(cmds["search children"], child)
	rows, _ := conn.Query(context.Background(), q)
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

func searchParents(conn *pgxpool.Pool, parent string) (parents []string) {
	q := fmt.Sprintf(cmds["search parents"], parent)
	rows, _ := conn.Query(context.Background(), q)
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

func SearchClasses(conn *pgxpool.Pool, class string)(classes []string){
	q := fmt.Sprintf(cmds["search classes"], class)
	rows, _ := conn.Query(context.Background(), q)
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

func SearchParentsAndChildren(conn *pgxpool.Pool, class string) (classes []string) {
	qp := fmt.Sprintf(cmds["search parents"], class)
	qc := fmt.Sprintf(cmds["search children"], class)
	q := fmt.Sprintf("select * from ((%v) union (%v)) as mixed", qp, qc)
	rows, _ := conn.Query(context.Background(), q)
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

func FuzzySearch(conn *pgxpool.Pool, module string, classes []string) (files []string) {
	// classes has mutiple c's, use c to get its similiar names in class_relation (as cc's), each cc has multiple children ccc
	condArray := []string{}
	for _, c := range classes {
		cc := SearchClasses(conn, c)
		// fmt.Printf("cc:%v\n", cc)
		for j, x := range cc {
			ccc := append(GetAllChildren(conn, x), x)
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
	rows, _ := conn.Query(context.Background(), q)
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

func GetAllFilesByClass(conn *pgxpool.Pool, class string) (files []string) {
	q := fmt.Sprintf(cmds["search files by class"], class) 
	rows, _ := conn.Query(context.Background(), q)
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
