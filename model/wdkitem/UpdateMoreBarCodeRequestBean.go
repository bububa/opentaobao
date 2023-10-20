package wdkitem

import (
	"sync"
)

// UpdateMoreBarCodeRequestBean 结构体
type UpdateMoreBarCodeRequestBean struct {
	// list
	BarCodeBoList []Barcodebolist `json:"bar_code_bo_list,omitempty" xml:"bar_code_bo_list>barcodebolist,omitempty"`
	// 机构编码
	OrgCode string `json:"org_code,omitempty" xml:"org_code,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
}

var poolUpdateMoreBarCodeRequestBean = sync.Pool{
	New: func() any {
		return new(UpdateMoreBarCodeRequestBean)
	},
}

// GetUpdateMoreBarCodeRequestBean() 从对象池中获取UpdateMoreBarCodeRequestBean
func GetUpdateMoreBarCodeRequestBean() *UpdateMoreBarCodeRequestBean {
	return poolUpdateMoreBarCodeRequestBean.Get().(*UpdateMoreBarCodeRequestBean)
}

// ReleaseUpdateMoreBarCodeRequestBean 释放UpdateMoreBarCodeRequestBean
func ReleaseUpdateMoreBarCodeRequestBean(v *UpdateMoreBarCodeRequestBean) {
	v.BarCodeBoList = v.BarCodeBoList[:0]
	v.OrgCode = ""
	v.SkuCode = ""
	poolUpdateMoreBarCodeRequestBean.Put(v)
}
