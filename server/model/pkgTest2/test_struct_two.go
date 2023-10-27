// 自动生成模板TestStructTwo
package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 测试结构二 结构体  TestStructTwo
type TestStructTwo struct {
	global.GVA_MODEL
	TestF string `json:"testF" form:"testF" gorm:"column:test_f;comment:;"` //测试字段
}

// TableName 测试结构二 TestStructTwo自定义表名 test_struct_two
func (TestStructTwo) TableName() string {
	return "test_struct_two"
}
