package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2"
	pkgTest2Req "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2/request"
)

type TestStructTwoService struct {
}

// CreateTestStructTwo 创建测试结构二记录
// Author [piexlmax](https://github.com/piexlmax)
func (tstService *TestStructTwoService) CreateTestStructTwo(tst *pkgTest2.TestStructTwo) (err error) {
	err = global.GVA_DB.Create(tst).Error
	return err
}

// DeleteTestStructTwo 删除测试结构二记录
// Author [piexlmax](https://github.com/piexlmax)
func (tstService *TestStructTwoService) DeleteTestStructTwo(tst pkgTest2.TestStructTwo) (err error) {
	err = global.GVA_DB.Delete(&tst).Error
	return err
}

// DeleteTestStructTwoByIds 批量删除测试结构二记录
// Author [piexlmax](https://github.com/piexlmax)
func (tstService *TestStructTwoService) DeleteTestStructTwoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest2.TestStructTwo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTestStructTwo 更新测试结构二记录
// Author [piexlmax](https://github.com/piexlmax)
func (tstService *TestStructTwoService) UpdateTestStructTwo(tst pkgTest2.TestStructTwo) (err error) {
	err = global.GVA_DB.Save(&tst).Error
	return err
}

// GetTestStructTwo 根据id获取测试结构二记录
// Author [piexlmax](https://github.com/piexlmax)
func (tstService *TestStructTwoService) GetTestStructTwo(id uint) (tst pkgTest2.TestStructTwo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tst).Error
	return
}

// GetTestStructTwoInfoList 分页获取测试结构二记录
// Author [piexlmax](https://github.com/piexlmax)
func (tstService *TestStructTwoService) GetTestStructTwoInfoList(info pkgTest2Req.TestStructTwoSearch) (list []pkgTest2.TestStructTwo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&pkgTest2.TestStructTwo{})
	var tsts []pkgTest2.TestStructTwo
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

	err = db.Find(&tsts).Error
	return tsts, total, err
}
