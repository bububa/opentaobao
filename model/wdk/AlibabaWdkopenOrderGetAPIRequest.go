package wdk

import (
	"net/url"
	"sync"

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
	// 外部主订单ID
	_outOrderId string
	// 五道口主订单id
	_bizOrderId int64
}

// NewAlibabaWdkopenOrderGetRequest 初始化AlibabaWdkopenOrderGetAPIRequest对象
func NewAlibabaWdkopenOrderGetRequest() *AlibabaWdkopenOrderGetAPIRequest {
	return &AlibabaWdkopenOrderGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkopenOrderGetAPIRequest) Reset() {
	r._storeId = ""
	r._outOrderId = ""
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkopenOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkopen.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkopenOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkopenOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 经营店id
func (r *AlibabaWdkopenOrderGetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkopenOrderGetAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetOutOrderId is OutOrderId Setter
// 外部主订单ID
func (r *AlibabaWdkopenOrderGetAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaWdkopenOrderGetAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetBizOrderId is BizOrderId Setter
// 五道口主订单id
func (r *AlibabaWdkopenOrderGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaWdkopenOrderGetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolAlibabaWdkopenOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkopenOrderGetRequest()
	},
}

// GetAlibabaWdkopenOrderGetRequest 从 sync.Pool 获取 AlibabaWdkopenOrderGetAPIRequest
func GetAlibabaWdkopenOrderGetAPIRequest() *AlibabaWdkopenOrderGetAPIRequest {
	return poolAlibabaWdkopenOrderGetAPIRequest.Get().(*AlibabaWdkopenOrderGetAPIRequest)
}

// ReleaseAlibabaWdkopenOrderGetAPIRequest 将 AlibabaWdkopenOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkopenOrderGetAPIRequest(v *AlibabaWdkopenOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkopenOrderGetAPIRequest.Put(v)
}
