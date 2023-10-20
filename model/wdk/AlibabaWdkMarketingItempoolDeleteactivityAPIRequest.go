package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempooldeleteactivityAPIRequest 删除商品池活动 API请求
// alibaba.wdk.marketing.itempool.deleteactivity
//
// 删除商品池活动
type AlibabawdkmarketingitempooldeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabawdkmarketingitempooldeleteactivityRequest 初始化AlibabawdkmarketingitempooldeleteactivityAPIRequest对象
func NewAlibabawdkmarketingitempooldeleteactivityRequest() *AlibabawdkmarketingitempooldeleteactivityAPIRequest {
	return &AlibabawdkmarketingitempooldeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempooldeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempooldeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempooldeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabawdkmarketingitempooldeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitempooldeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
