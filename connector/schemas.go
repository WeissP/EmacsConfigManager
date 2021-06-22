package connector

var (
	cmds = map[string]string{
		"create tables": `
  CREATE TABLE file (
    filename varchar NOT NULL PRIMARY KEY,
    module_name varchar NOT NULL,
    class varchar NOT NULL
  );` + `CREATE TABLE class_relation (
    child varchar NOT NULL PRIMARY KEY,
    parent varchar NOT NULL
  );`,
		"drop tables": `drop table if exists file;` + `drop table if exists class_relation;`,
		"add file": `INSERT INTO file(filename, module_name, class) values ('%s', '%s', '%s');`,
		"add parent": `INSERT INTO class_relation(child, parent) values ('%s', '%s');`,
		"get parent": `select parent from class_relation where child = '%s'`,
		"get child": `select child from class_relation where parent = '%s'`,
		"search children": `select child from class_relation where child like '%%%s%%'`,
		"search parents": `select parent from class_relation where parent like '%%%s%%'`,		
		"is child exists": `select child from class_relation where child = '%s'`,
		"is parent exists": `select parent from class_relation where parent = '%s'`,
		"search files by class": `select filename from file where class = '%s'`,
		"delete file": `delete from file where module_name = '%s' and class = '%s'`,
		"delete relation": `delete from class_relation where child = '%s'`,
		"search classes": `select distinct class from file where class like '%%%s%%'`,
	}
)
