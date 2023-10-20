package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolqueryactivityAPIRequest 查找商品池活动 API请求
// alibaba.wdk.marketing.itempool.queryactivity
//
// 查找商品池活动
type AlibabawdkmarketingitempoolqueryactivityAPIRequest struct {
	model.Params
	// 查询商品池活动入参
	_param0 *CommonActivityParam
}

// NewAlibabawdkmarketingitempoolqueryactivityRequest 初始化AlibabawdkmarketingitempoolqueryactivityAPIRequest对象
func NewAlibabawdkmarketingitempoolqueryactivityRequest() *AlibabawdkmarketingitempoolqueryactivityAPIRequest {
	return &AlibabawdkmarketingitempoolqueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempoolqueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempoolqueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempoolqueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询商品池活动入参
func (r *AlibabawdkmarketingitempoolqueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingitempoolqueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
