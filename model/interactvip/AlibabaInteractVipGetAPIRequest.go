package interactvip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractvipgetAPIRequest 会员淘气值获取 API请求
// alibaba.interact.vip.get
//
// 提供用户淘气值&amp;用户角色身份查询
type AlibabainteractvipgetAPIRequest struct {
	model.Params
}

// NewAlibabainteractvipgetRequest 初始化AlibabainteractvipgetAPIRequest对象
func NewAlibabainteractvipgetRequest() *AlibabainteractvipgetAPIRequest {
	return &AlibabainteractvipgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractvipgetAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.vip.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractvipgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractvipgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
