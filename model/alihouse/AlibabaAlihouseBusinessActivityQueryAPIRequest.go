package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseBusinessActivityQueryAPIRequest 电商券活动公司数据查询 API请求
// alibaba.alihouse.business.activity.query
//
// 电商券活动公司数据查询
type AlibabaAlihouseBusinessActivityQueryAPIRequest struct {
	model.Params
	// 公司主体ID
	_merchantOpenId int64
}

// NewAlibabaAlihouseBusinessActivityQueryRequest 初始化AlibabaAlihouseBusinessActivityQueryAPIRequest对象
func NewAlibabaAlihouseBusinessActivityQueryRequest() *AlibabaAlihouseBusinessActivityQueryAPIRequest {
	return &AlibabaAlihouseBusinessActivityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseBusinessActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.business.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseBusinessActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseBusinessActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantOpenId is MerchantOpenId Setter
// 公司主体ID
func (r *AlibabaAlihouseBusinessActivityQueryAPIRequest) SetMerchantOpenId(_merchantOpenId int64) error {
	r._merchantOpenId = _merchantOpenId
	r.Set("merchant_open_id", _merchantOpenId)
	return nil
}

// GetMerchantOpenId MerchantOpenId Getter
func (r AlibabaAlihouseBusinessActivityQueryAPIRequest) GetMerchantOpenId() int64 {
	return r._merchantOpenId
}
