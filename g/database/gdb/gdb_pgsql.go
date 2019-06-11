<<<<<<< HEAD
// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.
=======
// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
>>>>>>> upstream/master


package gdb

import (
    "fmt"
    "regexp"
    "database/sql"
)

<<<<<<< HEAD
// postgresql的适配
// @todo 需要完善replace和save的操作覆盖

// 数据库链接对象
type dbpgsql struct {
    Db
}

// 创建SQL操作对象，内部采用了lazy link处理
func (db *dbpgsql) Open (c *ConfigNode) (*sql.DB, error) {
    var source string
    if c.Linkinfo != "" {
        source = c.Linkinfo
    } else {
        source = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", c.User, c.Pass, c.Host, c.Port, c.Name)
=======
// PostgreSQL的适配.
// 使用时需要import:
// _ "github.com/gogf/gf/third/github.com/lib/pq"
// @todo 需要完善replace和save的操作覆盖

// 数据库链接对象
type dbPgsql struct {
    *dbBase
}

// 创建SQL操作对象，内部采用了lazy link处理
func (db *dbPgsql) Open (config *ConfigNode) (*sql.DB, error) {
    var source string
    if config.LinkInfo != "" {
        source = config.LinkInfo
    } else {
        source = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", config.User, config.Pass, config.Host, config.Port, config.Name)
>>>>>>> upstream/master
    }
    if db, err := sql.Open("postgres", source); err == nil {
        return db, nil
    } else {
        return nil, err
    }
}

<<<<<<< HEAD
// 获得关键字操作符 - 左
func (db *dbpgsql) getQuoteCharLeft () string {
    return "\""
}

// 获得关键字操作符 - 右
func (db *dbpgsql) getQuoteCharRight () string {
    return "\""
}

// 在执行sql之前对sql进行进一步处理
func (db *dbpgsql) handleSqlBeforeExec(q *string) *string {
    reg   := regexp.MustCompile("\\?")
    index := 0
    str   := reg.ReplaceAllStringFunc(*q, func (s string) string {
        index ++
        return fmt.Sprintf("$%d", index)
    })
    return &str
}

=======
// 获得关键字操作符
func (db *dbPgsql) getChars () (charLeft string, charRight string) {
    return "\"", "\""
}

// 在执行sql之前对sql进行进一步处理
func (db *dbPgsql) handleSqlBeforeExec(query string) string {
    reg   := regexp.MustCompile("\\?")
    index := 0
    str   := reg.ReplaceAllStringFunc(query, func (s string) string {
        index ++
        return fmt.Sprintf("$%d", index)
    })
    return str
}
>>>>>>> upstream/master
