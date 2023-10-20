package lstlogistics

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstShiporderCreateAPIRequest) Reset() {
	r._shipOrder = nil
	r.Params.ToZero()
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

var poolAlibabaLstShiporderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstShiporderCreateRequest()
	},
}

// GetAlibabaLstShiporderCreateRequest 从 sync.Pool 获取 AlibabaLstShiporderCreateAPIRequest
func GetAlibabaLstShiporderCreateAPIRequest() *AlibabaLstShiporderCreateAPIRequest {
	return poolAlibabaLstShiporderCreateAPIRequest.Get().(*AlibabaLstShiporderCreateAPIRequest)
}

// ReleaseAlibabaLstShiporderCreateAPIRequest 将 AlibabaLstShiporderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstShiporderCreateAPIRequest(v *AlibabaLstShiporderCreateAPIRequest) {
	v.Reset()
	poolAlibabaLstShiporderCreateAPIRequest.Put(v)
}
