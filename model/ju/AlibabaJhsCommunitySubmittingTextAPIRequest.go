package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajhscommunitysubmittingtextAPIRequest 聚划算社群动态文案下发接口 API请求
// alibaba.jhs.community.submitting.text
//
// 聚划算社群动态文案下发接口
type AlibabajhscommunitysubmittingtextAPIRequest struct {
	model.Params
}

// NewAlibabajhscommunitysubmittingtextRequest 初始化AlibabajhscommunitysubmittingtextAPIRequest对象
func NewAlibabajhscommunitysubmittingtextRequest() *AlibabajhscommunitysubmittingtextAPIRequest {
	return &AlibabajhscommunitysubmittingtextAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajhscommunitysubmittingtextAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.submitting.text"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajhscommunitysubmittingtextAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajhscommunitysubmittingtextAPIRequest) GetRawParams() model.Params {
	return r.Params
}
