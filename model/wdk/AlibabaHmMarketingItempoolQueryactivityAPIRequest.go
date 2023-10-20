package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolqueryactivityAPIRequest 查找商品池活动 API请求
// alibaba.hm.marketing.itempool.queryactivity
//
// 查找商品池活动
type AlibabahmmarketingitempoolqueryactivityAPIRequest struct {
	model.Params
	// 查询商品池活动入参
	_param0 *CommonActivityParam
}

// NewAlibabahmmarketingitempoolqueryactivityRequest 初始化AlibabahmmarketingitempoolqueryactivityAPIRequest对象
func NewAlibabahmmarketingitempoolqueryactivityRequest() *AlibabahmmarketingitempoolqueryactivityAPIRequest {
	return &AlibabahmmarketingitempoolqueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempoolqueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempoolqueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempoolqueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询商品池活动入参
func (r *AlibabahmmarketingitempoolqueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingitempoolqueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
