package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkopenOrderGetAPIRequest 五道口商户订单获取 API请求
// alibaba.wdkopen.order.get
//
// 商户通过五道口订单id获取订单信息
type AlibabaWdkopenOrderGetAPIRequest struct {
	model.Params
	// 经营店id
	_storeId string
	// 五道口主订单id
	_bizOrderId int64
	// 外部主订单ID
	_outOrderId string
}

// NewAlibabaWdkopenOrderGetRequest 初始化AlibabaWdkopenOrderGetAPIRequest对象
func NewAlibabaWdkopenOrderGetRequest() *AlibabaWdkopenOrderGetAPIRequest {
	return &AlibabaWdkopenOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkopenOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkopen.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkopenOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 经营店id
func (r *AlibabaWdkopenOrderGetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaWdkopenOrderGetAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is BizOrderId Setter
// 五道口主订单id
func (r *AlibabaWdkopenOrderGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// Get BizOrderId Getter
func (r AlibabaWdkopenOrderGetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// Set is OutOrderId Setter
// 外部主订单ID
func (r *AlibabaWdkopenOrderGetAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// Get OutOrderId Getter
func (r AlibabaWdkopenOrderGetAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}
