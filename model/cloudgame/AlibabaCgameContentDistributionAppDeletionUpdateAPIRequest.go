package cloudgame

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest) Reset() {
	r._reqParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.content.distribution.app.deletion.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaCgameContentDistributionAppDeletionUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCgameContentDistributionAppDeletionUpdateRequest()
	},
}

// GetAlibabaCgameContentDistributionAppDeletionUpdateRequest 从 sync.Pool 获取 AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest
func GetAlibabaCgameContentDistributionAppDeletionUpdateAPIRequest() *AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest {
	return poolAlibabaCgameContentDistributionAppDeletionUpdateAPIRequest.Get().(*AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest)
}

// ReleaseAlibabaCgameContentDistributionAppDeletionUpdateAPIRequest 将 AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaCgameContentDistributionAppDeletionUpdateAPIRequest(v *AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest) {
	v.Reset()
	poolAlibabaCgameContentDistributionAppDeletionUpdateAPIRequest.Put(v)
}
