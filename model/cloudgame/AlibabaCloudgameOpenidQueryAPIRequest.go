package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameopenidqueryAPIRequest 咖哒用户ID查询 API请求
// alibaba.cloudgame.openid.query
//
// 云游戏业务需要提供查询用户信息的TOP能力
type AlibabacloudgameopenidqueryAPIRequest struct {
	model.Params
	// 入参model
	_requestParam *HavanaUserIdQueryRequest
}

// NewAlibabacloudgameopenidqueryRequest 初始化AlibabacloudgameopenidqueryAPIRequest对象
func NewAlibabacloudgameopenidqueryRequest() *AlibabacloudgameopenidqueryAPIRequest {
	return &AlibabacloudgameopenidqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameopenidqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.openid.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameopenidqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameopenidqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestParam is RequestParam Setter
// 入参model
func (r *AlibabacloudgameopenidqueryAPIRequest) SetRequestParam(_requestParam *HavanaUserIdQueryRequest) error {
	r._requestParam = _requestParam
	r.Set("request_param", _requestParam)
	return nil
}

// GetRequestParam RequestParam Getter
func (r AlibabacloudgameopenidqueryAPIRequest) GetRequestParam() *HavanaUserIdQueryRequest {
	return r._requestParam
}
