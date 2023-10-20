package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgamecontentdistributionappdeletionupdateAPIRequest 游戏删除回调 API请求
// alibaba.cgame.content.distribution.app.deletion.update
//
// 游戏删除回调
type AlibabacgamecontentdistributionappdeletionupdateAPIRequest struct {
	model.Params
	// 请求参数
	_reqParam *AppDeletionCallbackRequest
}

// NewAlibabacgamecontentdistributionappdeletionupdateRequest 初始化AlibabacgamecontentdistributionappdeletionupdateAPIRequest对象
func NewAlibabacgamecontentdistributionappdeletionupdateRequest() *AlibabacgamecontentdistributionappdeletionupdateAPIRequest {
	return &AlibabacgamecontentdistributionappdeletionupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacgamecontentdistributionappdeletionupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.content.distribution.app.deletion.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacgamecontentdistributionappdeletionupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacgamecontentdistributionappdeletionupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqParam is ReqParam Setter
// 请求参数
func (r *AlibabacgamecontentdistributionappdeletionupdateAPIRequest) SetReqParam(_reqParam *AppDeletionCallbackRequest) error {
	r._reqParam = _reqParam
	r.Set("req_param", _reqParam)
	return nil
}

// GetReqParam ReqParam Getter
func (r AlibabacgamecontentdistributionappdeletionupdateAPIRequest) GetReqParam() *AppDeletionCallbackRequest {
	return r._reqParam
}
