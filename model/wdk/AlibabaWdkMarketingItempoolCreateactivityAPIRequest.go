package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolcreateactivityAPIRequest 添加商品池活动 API请求
// alibaba.wdk.marketing.itempool.createactivity
//
// 添加商品池活动
type AlibabawdkmarketingitempoolcreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// NewAlibabawdkmarketingitempoolcreateactivityRequest 初始化AlibabawdkmarketingitempoolcreateactivityAPIRequest对象
func NewAlibabawdkmarketingitempoolcreateactivityRequest() *AlibabawdkmarketingitempoolcreateactivityAPIRequest {
	return &AlibabawdkmarketingitempoolcreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempoolcreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempoolcreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempoolcreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabawdkmarketingitempoolcreateactivityAPIRequest) SetParam(_param *ItemPoolActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitempoolcreateactivityAPIRequest) GetParam() *ItemPoolActivity {
	return r._param
}
