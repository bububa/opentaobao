package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitemdiscountcreateactivityAPIRequest 创建商品特价活动 API请求
// alibaba.hm.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
type AlibabahmmarketingitemdiscountcreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemDiscountActivityRequest
}

// NewAlibabahmmarketingitemdiscountcreateactivityRequest 初始化AlibabahmmarketingitemdiscountcreateactivityAPIRequest对象
func NewAlibabahmmarketingitemdiscountcreateactivityRequest() *AlibabahmmarketingitemdiscountcreateactivityAPIRequest {
	return &AlibabahmmarketingitemdiscountcreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitemdiscountcreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitemdiscountcreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitemdiscountcreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabahmmarketingitemdiscountcreateactivityAPIRequest) SetParam(_param *ItemDiscountActivityRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitemdiscountcreateactivityAPIRequest) GetParam() *ItemDiscountActivityRequest {
	return r._param
}
