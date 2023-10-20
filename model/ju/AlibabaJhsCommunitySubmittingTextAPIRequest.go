package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunitySubmittingTextAPIRequest 聚划算社群动态文案下发接口 API请求
// alibaba.jhs.community.submitting.text
//
// 聚划算社群动态文案下发接口
type AlibabaJhsCommunitySubmittingTextAPIRequest struct {
	model.Params
}

// NewAlibabaJhsCommunitySubmittingTextRequest 初始化AlibabaJhsCommunitySubmittingTextAPIRequest对象
func NewAlibabaJhsCommunitySubmittingTextRequest() *AlibabaJhsCommunitySubmittingTextAPIRequest {
	return &AlibabaJhsCommunitySubmittingTextAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJhsCommunitySubmittingTextAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.submitting.text"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJhsCommunitySubmittingTextAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJhsCommunitySubmittingTextAPIRequest) GetRawParams() model.Params {
	return r.Params
}
