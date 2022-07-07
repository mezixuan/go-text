package test

import (
	"bytes"
	"cloud-disk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	//dataSourceName  root用户密码@/数据库名?charset=utf8mb4
	engin, err := xorm.NewEngine("mysql", "root:123456@/cloud-disk?charset=utf8mb4")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	err = engin.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}
