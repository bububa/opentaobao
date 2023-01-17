package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivitySellerinfoAPIRequest 商家信息推送 API请求
// alibaba.lsy.crm.activity.sellerinfo
//
// 本地团商家信息推送
type AlibabaLsyCrmActivitySellerinfoAPIRequest struct {
	model.Params
}

// NewAlibabaLsyCrmActivitySellerinfoRequest 初始化AlibabaLsyCrmActivitySellerinfoAPIRequest对象
func NewAlibabaLsyCrmActivitySellerinfoRequest() *AlibabaLsyCrmActivitySellerinfoAPIRequest {
	return &AlibabaLsyCrmActivitySellerinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivitySellerinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.sellerinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivitySellerinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyCrmActivitySellerinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}
