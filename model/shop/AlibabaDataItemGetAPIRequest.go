package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDataItemGetAPIRequest 获取商品 API请求
// alibaba.data.item.get
//
// 获取商品信息，作为客户端Weex鉴权的虚拟api
type AlibabaDataItemGetAPIRequest struct {
	model.Params
	// 获取商品信息，作为客户端Weex鉴权的虚拟api
	_unNamed string
}

// NewAlibabaDataItemGetRequest 初始化AlibabaDataItemGetAPIRequest对象
func NewAlibabaDataItemGetRequest() *AlibabaDataItemGetAPIRequest {
	return &AlibabaDataItemGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDataItemGetAPIRequest) Reset() {
	r._unNamed = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDataItemGetAPIRequest) GetApiMethodName() string {
	return "alibaba.data.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDataItemGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDataItemGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnNamed is UnNamed Setter
// 获取商品信息，作为客户端Weex鉴权的虚拟api
func (r *AlibabaDataItemGetAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// GetUnNamed UnNamed Getter
func (r AlibabaDataItemGetAPIRequest) GetUnNamed() string {
	return r._unNamed
}

var poolAlibabaDataItemGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDataItemGetRequest()
	},
}

// GetAlibabaDataItemGetRequest 从 sync.Pool 获取 AlibabaDataItemGetAPIRequest
func GetAlibabaDataItemGetAPIRequest() *AlibabaDataItemGetAPIRequest {
	return poolAlibabaDataItemGetAPIRequest.Get().(*AlibabaDataItemGetAPIRequest)
}

// ReleaseAlibabaDataItemGetAPIRequest 将 AlibabaDataItemGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaDataItemGetAPIRequest(v *AlibabaDataItemGetAPIRequest) {
	v.Reset()
	poolAlibabaDataItemGetAPIRequest.Put(v)
}
