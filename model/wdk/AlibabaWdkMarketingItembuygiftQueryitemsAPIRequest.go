package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitembuygiftqueryitemsAPIRequest 查询买赠活动下的商品 API请求
// alibaba.wdk.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
type AlibabawdkmarketingitembuygiftqueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabawdkmarketingitembuygiftqueryitemsRequest 初始化AlibabawdkmarketingitembuygiftqueryitemsAPIRequest对象
func NewAlibabawdkmarketingitembuygiftqueryitemsRequest() *AlibabawdkmarketingitembuygiftqueryitemsAPIRequest {
	return &AlibabawdkmarketingitembuygiftqueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitembuygiftqueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitembuygiftqueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitembuygiftqueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabawdkmarketingitembuygiftqueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitembuygiftqueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
