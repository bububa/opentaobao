package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstShiporderCreateAPIRequest 零售通发货单创建 API请求
// alibaba.lst.shiporder.create
//
// 通过该接口可以创建零售通运保保发货单，并处理相关业务流程。
type AlibabaLstShiporderCreateAPIRequest struct {
	model.Params
	// 创建发货单入参
	_shipOrder *LstThirdPartMainShipOrderCreateDto
}

// NewAlibabaLstShiporderCreateRequest 初始化AlibabaLstShiporderCreateAPIRequest对象
func NewAlibabaLstShiporderCreateRequest() *AlibabaLstShiporderCreateAPIRequest {
	return &AlibabaLstShiporderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstShiporderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.shiporder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstShiporderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstShiporderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShipOrder is ShipOrder Setter
// 创建发货单入参
func (r *AlibabaLstShiporderCreateAPIRequest) SetShipOrder(_shipOrder *LstThirdPartMainShipOrderCreateDto) error {
	r._shipOrder = _shipOrder
	r.Set("ship_order", _shipOrder)
	return nil
}

// GetShipOrder ShipOrder Getter
func (r AlibabaLstShiporderCreateAPIRequest) GetShipOrder() *LstThirdPartMainShipOrderCreateDto {
	return r._shipOrder
}
