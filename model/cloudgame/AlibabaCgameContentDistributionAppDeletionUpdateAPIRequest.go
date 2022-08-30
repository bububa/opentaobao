package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest 游戏删除回调 API请求
// alibaba.cgame.content.distribution.app.deletion.update
//
// 游戏删除回调
type AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_reqParam *AppDeletionCallbackRequest
}

// NewAlibabaCgameContentDistributionAppDeletionUpdateRequest 初始化AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest对象
func NewAlibabaCgameContentDistributionAppDeletionUpdateRequest() *AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest {
	return &AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.content.distribution.app.deletion.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReqParam is ReqParam Setter
// 请求参数
func (r *AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest) SetReqParam(_reqParam *AppDeletionCallbackRequest) error {
	r._reqParam = _reqParam
	r.Set("req_param", _reqParam)
	return nil
}

// GetReqParam ReqParam Getter
func (r AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest) GetReqParam() *AppDeletionCallbackRequest {
	return r._reqParam
}
