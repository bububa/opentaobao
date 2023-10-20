package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendingequipmentqueryAPIRequest 自动售卖机设备信息查询 API请求
// alibaba.lst.vending.equipment.query
//
// 为零售通品牌商提供已租赁的设备信息查询。
type AlibabalstvendingequipmentqueryAPIRequest struct {
	model.Params
	// 设备查询条件
	_openEquipmentQuery *OpenEquipmentQuery
}

// NewAlibabalstvendingequipmentqueryRequest 初始化AlibabalstvendingequipmentqueryAPIRequest对象
func NewAlibabalstvendingequipmentqueryRequest() *AlibabalstvendingequipmentqueryAPIRequest {
	return &AlibabalstvendingequipmentqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstvendingequipmentqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.equipment.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstvendingequipmentqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstvendingequipmentqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenEquipmentQuery is OpenEquipmentQuery Setter
// 设备查询条件
func (r *AlibabalstvendingequipmentqueryAPIRequest) SetOpenEquipmentQuery(_openEquipmentQuery *OpenEquipmentQuery) error {
	r._openEquipmentQuery = _openEquipmentQuery
	r.Set("open_equipment_query", _openEquipmentQuery)
	return nil
}

// GetOpenEquipmentQuery OpenEquipmentQuery Getter
func (r AlibabalstvendingequipmentqueryAPIRequest) GetOpenEquipmentQuery() *OpenEquipmentQuery {
	return r._openEquipmentQuery
}
