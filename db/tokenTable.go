package db

import (
	"fmt"
)

func QueryAll(table string) (map[string]string, error) {
	values := make(map[string]string)
	s := fmt.Sprintf("SELECT * FROM %s", table)
	if rows, err := sdb.Query(s); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var uid int
			var name, value string
			if err := rows.Scan(&uid, &name, &value); err != nil {
				continue
			}
			values[name] = value
		}
	}
	return values, nil
}

func Query(table, name string) (string, error) {
	s := fmt.Sprintf("SELECT * FROM %s WHERE name='%s'", table, name)
	if rows, err := sdb.Query(s); err != nil {
		return "", err
	} else {
		defer rows.Close()
		for rows.Next() {
			var uid int
			var name, value string
			if err := rows.Scan(&uid, &name, &value); err != nil {
				continue
			}
			return value, nil
		}
	}
	return "", nil
}

func Add(table, name, value string) error {
	s := fmt.Sprintf("INSERT INTO %s(name, value) values(?, ?)", table)
	if stmt, err := sdb.Prepare(s); err != nil {
		return err
	} else {
		if _, err := stmt.Exec(name, value); err != nil {
			return err
		}
	}
	return nil
}

func Del(table, name string) error {
	s := fmt.Sprintf("DELETE FROM %s where name=?", table)
	if stmt, err := sdb.Prepare(s); err != nil {
		return err
	} else {
		if _, err := stmt.Exec(name); err != nil {
			return err
		}
	}
	return nil
}

func Update(table, name, value string) error {
	if qv, err := Query(table, name); err == nil {
		if qv == "" {
			return Add(table, name, value)
		} else {
			return update(table, name, value)
		}
	} else {
		return err
	}
}

func update(table, name, value string) error {
	s := fmt.Sprintf("UPDATE %s set value=? where name=?", table)
	if stmt, err := sdb.Prepare(s); err != nil {
		return err
	} else {
		if _, err := stmt.Exec(value, name); err != nil {
			return err
		}
	}
	return nil
}
