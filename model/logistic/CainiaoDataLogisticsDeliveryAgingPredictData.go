package logistic

import (
	"sync"
)

// CainiaoDataLogisticsDeliveryAgingPredictData 结构体
type CainiaoDataLogisticsDeliveryAgingPredictData struct {
	// 快递公司id
	CompanyId string `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 配送时效。字符串，ISV拿到直接展示，不要做处理
	DeliveryPeriod string `json:"delivery_period,omitempty" xml:"delivery_period,omitempty"`
}

var poolCainiaoDataLogisticsDeliveryAgingPredictData = sync.Pool{
	New: func() any {
		return new(CainiaoDataLogisticsDeliveryAgingPredictData)
	},
}

// GetCainiaoDataLogisticsDeliveryAgingPredictData() 从对象池中获取CainiaoDataLogisticsDeliveryAgingPredictData
func GetCainiaoDataLogisticsDeliveryAgingPredictData() *CainiaoDataLogisticsDeliveryAgingPredictData {
	return poolCainiaoDataLogisticsDeliveryAgingPredictData.Get().(*CainiaoDataLogisticsDeliveryAgingPredictData)
}

// ReleaseCainiaoDataLogisticsDeliveryAgingPredictData 释放CainiaoDataLogisticsDeliveryAgingPredictData
func ReleaseCainiaoDataLogisticsDeliveryAgingPredictData(v *CainiaoDataLogisticsDeliveryAgingPredictData) {
	v.CompanyId = ""
	v.DeliveryPeriod = ""
	poolCainiaoDataLogisticsDeliveryAgingPredictData.Put(v)
}
