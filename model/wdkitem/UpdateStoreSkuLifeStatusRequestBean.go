package wdkitem

import (
	"sync"
)

// UpdateStoreSkuLifeStatusRequestBean 结构体
type UpdateStoreSkuLifeStatusRequestBean struct {
	// 机构编码
	OrgCode string `json:"org_code,omitempty" xml:"org_code,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 商品状态
	LifeStatus string `json:"life_status,omitempty" xml:"life_status,omitempty"`
	// 渠道编码
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 淘鲜达半日达项目新增，0 表示上架，1表示下架，当更新淘鲜达半日达渠道品上下架时，该字段必传
	OnlineSaleFlag int64 `json:"online_sale_flag,omitempty" xml:"online_sale_flag,omitempty"`
}

var poolUpdateStoreSkuLifeStatusRequestBean = sync.Pool{
	New: func() any {
		return new(UpdateStoreSkuLifeStatusRequestBean)
	},
}

// GetUpdateStoreSkuLifeStatusRequestBean() 从对象池中获取UpdateStoreSkuLifeStatusRequestBean
func GetUpdateStoreSkuLifeStatusRequestBean() *UpdateStoreSkuLifeStatusRequestBean {
	return poolUpdateStoreSkuLifeStatusRequestBean.Get().(*UpdateStoreSkuLifeStatusRequestBean)
}

// ReleaseUpdateStoreSkuLifeStatusRequestBean 释放UpdateStoreSkuLifeStatusRequestBean
func ReleaseUpdateStoreSkuLifeStatusRequestBean(v *UpdateStoreSkuLifeStatusRequestBean) {
	v.OrgCode = ""
	v.SkuCode = ""
	v.StoreId = ""
	v.LifeStatus = ""
	v.ShopId = ""
	v.OnlineSaleFlag = 0
	poolUpdateStoreSkuLifeStatusRequestBean.Put(v)
}
