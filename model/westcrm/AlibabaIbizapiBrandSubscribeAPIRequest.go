package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbizapiBrandSubscribeAPIRequest 关注品牌号 API请求
// alibaba.ibizapi.brand.subscribe
//
// 关注品牌号服务
type AlibabaIbizapiBrandSubscribeAPIRequest struct {
	model.Params
}

// NewAlibabaIbizapiBrandSubscribeRequest 初始化AlibabaIbizapiBrandSubscribeAPIRequest对象
func NewAlibabaIbizapiBrandSubscribeRequest() *AlibabaIbizapiBrandSubscribeAPIRequest {
	return &AlibabaIbizapiBrandSubscribeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIbizapiBrandSubscribeAPIRequest) GetApiMethodName() string {
	return "alibaba.ibizapi.brand.subscribe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIbizapiBrandSubscribeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
