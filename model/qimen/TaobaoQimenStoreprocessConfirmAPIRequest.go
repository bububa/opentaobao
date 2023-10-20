package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStoreprocessConfirmAPIRequest 仓内加工单确认接口 API请求
// taobao.qimen.storeprocess.confirm
//
// WMS调用奇门的接口,回传仓内加工单创建情况
type TaobaoQimenStoreprocessConfirmAPIRequest struct {
	model.Params
	//
	_request *StoreProcessConfirmRequest
}

// NewTaobaoQimenStoreprocessConfirmRequest 初始化TaobaoQimenStoreprocessConfirmAPIRequest对象
func NewTaobaoQimenStoreprocessConfirmRequest() *TaobaoQimenStoreprocessConfirmAPIRequest {
	return &TaobaoQimenStoreprocessConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenStoreprocessConfirmAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreprocessConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.storeprocess.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreprocessConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenStoreprocessConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenStoreprocessConfirmAPIRequest) SetRequest(_request *StoreProcessConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenStoreprocessConfirmAPIRequest) GetRequest() *StoreProcessConfirmRequest {
	return r._request
}

var poolTaobaoQimenStoreprocessConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenStoreprocessConfirmRequest()
	},
}

// GetTaobaoQimenStoreprocessConfirmRequest 从 sync.Pool 获取 TaobaoQimenStoreprocessConfirmAPIRequest
func GetTaobaoQimenStoreprocessConfirmAPIRequest() *TaobaoQimenStoreprocessConfirmAPIRequest {
	return poolTaobaoQimenStoreprocessConfirmAPIRequest.Get().(*TaobaoQimenStoreprocessConfirmAPIRequest)
}

// ReleaseTaobaoQimenStoreprocessConfirmAPIRequest 将 TaobaoQimenStoreprocessConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenStoreprocessConfirmAPIRequest(v *TaobaoQimenStoreprocessConfirmAPIRequest) {
	v.Reset()
	poolTaobaoQimenStoreprocessConfirmAPIRequest.Put(v)
}
