package lstvending

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingEquipmentQueryAPIRequest 自动售卖机设备信息查询 API请求
// alibaba.lst.vending.equipment.query
//
// 为零售通品牌商提供已租赁的设备信息查询。
type AlibabaLstVendingEquipmentQueryAPIRequest struct {
	model.Params
	// 设备查询条件
	_openEquipmentQuery *OpenEquipmentQuery
}

// NewAlibabaLstVendingEquipmentQueryRequest 初始化AlibabaLstVendingEquipmentQueryAPIRequest对象
func NewAlibabaLstVendingEquipmentQueryRequest() *AlibabaLstVendingEquipmentQueryAPIRequest {
	return &AlibabaLstVendingEquipmentQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstVendingEquipmentQueryAPIRequest) Reset() {
	r._openEquipmentQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingEquipmentQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.equipment.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingEquipmentQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstVendingEquipmentQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenEquipmentQuery is OpenEquipmentQuery Setter
// 设备查询条件
func (r *AlibabaLstVendingEquipmentQueryAPIRequest) SetOpenEquipmentQuery(_openEquipmentQuery *OpenEquipmentQuery) error {
	r._openEquipmentQuery = _openEquipmentQuery
	r.Set("open_equipment_query", _openEquipmentQuery)
	return nil
}

// GetOpenEquipmentQuery OpenEquipmentQuery Getter
func (r AlibabaLstVendingEquipmentQueryAPIRequest) GetOpenEquipmentQuery() *OpenEquipmentQuery {
	return r._openEquipmentQuery
}

var poolAlibabaLstVendingEquipmentQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstVendingEquipmentQueryRequest()
	},
}

// GetAlibabaLstVendingEquipmentQueryRequest 从 sync.Pool 获取 AlibabaLstVendingEquipmentQueryAPIRequest
func GetAlibabaLstVendingEquipmentQueryAPIRequest() *AlibabaLstVendingEquipmentQueryAPIRequest {
	return poolAlibabaLstVendingEquipmentQueryAPIRequest.Get().(*AlibabaLstVendingEquipmentQueryAPIRequest)
}

// ReleaseAlibabaLstVendingEquipmentQueryAPIRequest 将 AlibabaLstVendingEquipmentQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstVendingEquipmentQueryAPIRequest(v *AlibabaLstVendingEquipmentQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstVendingEquipmentQueryAPIRequest.Put(v)
}
