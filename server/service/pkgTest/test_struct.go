package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
)

type TestStructService struct {
}

// CreateTestStruct 创建测试结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (tsService *TestStructService) CreateTestStruct(ts *pkgTest.TestStruct) (err error) {
	err = global.GVA_DB.Create(ts).Error
	return err
}

// DeleteTestStruct 删除测试结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (tsService *TestStructService) DeleteTestStruct(ts pkgTest.TestStruct) (err error) {
	err = global.GVA_DB.Delete(&ts).Error
	return err
}

// DeleteTestStructByIds 批量删除测试结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (tsService *TestStructService) DeleteTestStructByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest.TestStruct{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTestStruct 更新测试结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (tsService *TestStructService) UpdateTestStruct(ts pkgTest.TestStruct) (err error) {
	err = global.GVA_DB.Save(&ts).Error
	return err
}

// GetTestStruct 根据id获取测试结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (tsService *TestStructService) GetTestStruct(id uint) (ts pkgTest.TestStruct, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ts).Error
	return
}

// GetTestStructInfoList 分页获取测试结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (tsService *TestStructService) GetTestStructInfoList(info pkgTestReq.TestStructSearch) (list []pkgTest.TestStruct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&pkgTest.TestStruct{})
	var tss []pkgTest.TestStruct
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tss).Error
	return tss, total, err
}
