package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaModifyaddressIsvBindsellerCheckAPIRequest 查询服务商下的商家是否开通了改地址 API请求
// alibaba.modifyaddress.isv.bindseller.check
//
// 鉴权服务商和商家的绑定关系，看商家是否开通了改地址
// 1. 没有授权
// 2. 已与当前appkey签约
// 3. 没有签约
// 4. 与其他服务商软件签约，如果是同一个isv name，返回appkey，否则不返回。
type AlibabaModifyaddressIsvBindsellerCheckAPIRequest struct {
	model.Params
}

// NewAlibabaModifyaddressIsvBindsellerCheckRequest 初始化AlibabaModifyaddressIsvBindsellerCheckAPIRequest对象
func NewAlibabaModifyaddressIsvBindsellerCheckRequest() *AlibabaModifyaddressIsvBindsellerCheckAPIRequest {
	return &AlibabaModifyaddressIsvBindsellerCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaModifyaddressIsvBindsellerCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.modifyaddress.isv.bindseller.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaModifyaddressIsvBindsellerCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaModifyaddressIsvBindsellerCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}
