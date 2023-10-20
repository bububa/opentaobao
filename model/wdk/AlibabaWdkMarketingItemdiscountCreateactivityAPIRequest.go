package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitemdiscountcreateactivityAPIRequest 创建商品特价活动 API请求
// alibaba.wdk.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
type AlibabawdkmarketingitemdiscountcreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemDiscountActivityRequest
}

// NewAlibabawdkmarketingitemdiscountcreateactivityRequest 初始化AlibabawdkmarketingitemdiscountcreateactivityAPIRequest对象
func NewAlibabawdkmarketingitemdiscountcreateactivityRequest() *AlibabawdkmarketingitemdiscountcreateactivityAPIRequest {
	return &AlibabawdkmarketingitemdiscountcreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitemdiscountcreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitemdiscountcreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitemdiscountcreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabawdkmarketingitemdiscountcreateactivityAPIRequest) SetParam(_param *ItemDiscountActivityRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitemdiscountcreateactivityAPIRequest) GetParam() *ItemDiscountActivityRequest {
	return r._param
}
