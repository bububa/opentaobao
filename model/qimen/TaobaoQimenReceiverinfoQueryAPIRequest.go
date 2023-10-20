package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReceiverinfoQueryAPIRequest OAID 收件人信息解密接口 API请求
// taobao.qimen.receiverinfo.query
//
// WMS 调用该接口，通过 OAID 查询解密后的收件人信息
type TaobaoQimenReceiverinfoQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenReceiverinfoQueryRequest
}

// NewTaobaoQimenReceiverinfoQueryRequest 初始化TaobaoQimenReceiverinfoQueryAPIRequest对象
func NewTaobaoQimenReceiverinfoQueryRequest() *TaobaoQimenReceiverinfoQueryAPIRequest {
	return &TaobaoQimenReceiverinfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenReceiverinfoQueryAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenReceiverinfoQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.receiverinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenReceiverinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenReceiverinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenReceiverinfoQueryAPIRequest) SetRequest(_request *TaobaoQimenReceiverinfoQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenReceiverinfoQueryAPIRequest) GetRequest() *TaobaoQimenReceiverinfoQueryRequest {
	return r._request
}

var poolTaobaoQimenReceiverinfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenReceiverinfoQueryRequest()
	},
}

// GetTaobaoQimenReceiverinfoQueryRequest 从 sync.Pool 获取 TaobaoQimenReceiverinfoQueryAPIRequest
func GetTaobaoQimenReceiverinfoQueryAPIRequest() *TaobaoQimenReceiverinfoQueryAPIRequest {
	return poolTaobaoQimenReceiverinfoQueryAPIRequest.Get().(*TaobaoQimenReceiverinfoQueryAPIRequest)
}

// ReleaseTaobaoQimenReceiverinfoQueryAPIRequest 将 TaobaoQimenReceiverinfoQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenReceiverinfoQueryAPIRequest(v *TaobaoQimenReceiverinfoQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenReceiverinfoQueryAPIRequest.Put(v)
}
