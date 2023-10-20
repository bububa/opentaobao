package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolactivitycreateAPIRequest 创建活动新接口 API请求
// alibaba.wdk.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
type AlibabawdkmarketingitempoolactivitycreateAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// NewAlibabawdkmarketingitempoolactivitycreateRequest 初始化AlibabawdkmarketingitempoolactivitycreateAPIRequest对象
func NewAlibabawdkmarketingitempoolactivitycreateRequest() *AlibabawdkmarketingitempoolactivitycreateAPIRequest {
	return &AlibabawdkmarketingitempoolactivitycreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempoolactivitycreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempoolactivitycreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempoolactivitycreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabawdkmarketingitempoolactivitycreateAPIRequest) SetParam(_param *ItemPoolActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitempoolactivitycreateAPIRequest) GetParam() *ItemPoolActivity {
	return r._param
}
