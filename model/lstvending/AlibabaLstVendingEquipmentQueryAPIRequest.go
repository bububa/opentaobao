package lstvending

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingEquipmentQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.equipment.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingEquipmentQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
