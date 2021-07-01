package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDataItemGetAPIRequest
获取商品 API请求
alibaba.data.item.get

获取商品信息，作为客户端Weex鉴权的虚拟api */
type AlibabaDataItemGetAPIRequest struct {
	model.Params
	// 获取商品信息，作为客户端Weex鉴权的虚拟api
	_unNamed string
}

// NewAlibabaDataItemGetRequest 初始化AlibabaDataItemGetAPIRequest对象
func NewAlibabaDataItemGetRequest() *AlibabaDataItemGetAPIRequest {
	return &AlibabaDataItemGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDataItemGetAPIRequest) GetApiMethodName() string {
	return "alibaba.data.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDataItemGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UnNamed Setter
// 获取商品信息，作为客户端Weex鉴权的虚拟api
func (r *AlibabaDataItemGetAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// Get UnNamed Getter
func (r AlibabaDataItemGetAPIRequest) GetUnNamed() string {
	return r._unNamed
}
