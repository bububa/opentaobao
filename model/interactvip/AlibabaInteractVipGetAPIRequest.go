package interactvip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractVipGetAPIRequest 会员淘气值获取 API请求
// alibaba.interact.vip.get
//
// 提供用户淘气值&amp;用户角色身份查询
type AlibabaInteractVipGetAPIRequest struct {
	model.Params
}

// NewAlibabaInteractVipGetRequest 初始化AlibabaInteractVipGetAPIRequest对象
func NewAlibabaInteractVipGetRequest() *AlibabaInteractVipGetAPIRequest {
	return &AlibabaInteractVipGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractVipGetAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.vip.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractVipGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractVipGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
