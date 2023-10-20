package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajhscommunityactivitylistAPIRequest 聚划算用增淘外社群服务活动列表 API请求
// alibaba.jhs.community.activity.list
//
// 聚划算用增淘外社群服务活动列表
type AlibabajhscommunityactivitylistAPIRequest struct {
	model.Params
}

// NewAlibabajhscommunityactivitylistRequest 初始化AlibabajhscommunityactivitylistAPIRequest对象
func NewAlibabajhscommunityactivitylistRequest() *AlibabajhscommunityactivitylistAPIRequest {
	return &AlibabajhscommunityactivitylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajhscommunityactivitylistAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.activity.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajhscommunityactivitylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajhscommunityactivitylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}
