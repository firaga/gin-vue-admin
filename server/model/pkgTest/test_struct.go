// 自动生成模板TestStruct
package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 测试结构 结构体  TestStruct
type TestStruct struct {
	global.GVA_MODEL
	TestF string `json:"testF" form:"testF" gorm:"column:test_f;comment:;"` //测试
}

// TableName 测试结构 TestStruct自定义表名 test_struct
func (TestStruct) TableName() string {
	return "test_struct"
}
