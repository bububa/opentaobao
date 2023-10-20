package omniorder

import (
	"sync"
)

// StoreConsignedResult 结构体
type StoreConsignedResult struct {
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 0表示无系统异常
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 物流公司名称
	LogisticCompany string `json:"logistic_company,omitempty" xml:"logistic_company,omitempty"`
	// 物流公司code，如果id和code都填入，以code为准。点点送：DISTRIBUTOR_12006531；门店自送：DISTRIBUTOR_12709653；如果是菜鸟配送，code和company可以为空
	LogisticCompanyCode string `json:"logistic_company_code,omitempty" xml:"logistic_company_code,omitempty"`
	// 物流公司id
	LogisticId string `json:"logistic_id,omitempty" xml:"logistic_id,omitempty"`
	// 物流单号
	LogisticNo string `json:"logistic_no,omitempty" xml:"logistic_no,omitempty"`
	// 异常描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作者
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 店铺Id, 可能是门店或者电商仓
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 店铺名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 店铺类型, 门店(Store)或者电商仓(Warehouse)
	StoreType string `json:"store_type,omitempty" xml:"store_type,omitempty"`
	// 子订单Id
	SubOid int64 `json:"sub_oid,omitempty" xml:"sub_oid,omitempty"`
	// 速店通packageId
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 主订单Id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolStoreConsignedResult = sync.Pool{
	New: func() any {
		return new(StoreConsignedResult)
	},
}

// GetStoreConsignedResult() 从对象池中获取StoreConsignedResult
func GetStoreConsignedResult() *StoreConsignedResult {
	return poolStoreConsignedResult.Get().(*StoreConsignedResult)
}

// ReleaseStoreConsignedResult 释放StoreConsignedResult
func ReleaseStoreConsignedResult(v *StoreConsignedResult) {
	v.Attributes = ""
	v.Code = ""
	v.LogisticCompany = ""
	v.LogisticCompanyCode = ""
	v.LogisticId = ""
	v.LogisticNo = ""
	v.Message = ""
	v.Operator = ""
	v.StoreId = ""
	v.StoreName = ""
	v.StoreType = ""
	v.SubOid = 0
	v.PackageId = 0
	v.Tid = 0
	poolStoreConsignedResult.Put(v)
}
