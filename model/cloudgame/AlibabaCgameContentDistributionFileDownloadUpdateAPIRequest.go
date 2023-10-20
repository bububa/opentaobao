package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgamecontentdistributionfiledownloadupdateAPIRequest 文件下载回调 API请求
// alibaba.cgame.content.distribution.file.download.update
//
// 提供内容服务器更新文件下载状态的能力
type AlibabacgamecontentdistributionfiledownloadupdateAPIRequest struct {
	model.Params
	// 请求参数
	_reqParam *FileDownloadCallbackRequest
}

// NewAlibabacgamecontentdistributionfiledownloadupdateRequest 初始化AlibabacgamecontentdistributionfiledownloadupdateAPIRequest对象
func NewAlibabacgamecontentdistributionfiledownloadupdateRequest() *AlibabacgamecontentdistributionfiledownloadupdateAPIRequest {
	return &AlibabacgamecontentdistributionfiledownloadupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacgamecontentdistributionfiledownloadupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.content.distribution.file.download.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacgamecontentdistributionfiledownloadupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacgamecontentdistributionfiledownloadupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqParam is ReqParam Setter
// 请求参数
func (r *AlibabacgamecontentdistributionfiledownloadupdateAPIRequest) SetReqParam(_reqParam *FileDownloadCallbackRequest) error {
	r._reqParam = _reqParam
	r.Set("req_param", _reqParam)
	return nil
}

// GetReqParam ReqParam Getter
func (r AlibabacgamecontentdistributionfiledownloadupdateAPIRequest) GetReqParam() *FileDownloadCallbackRequest {
	return r._reqParam
}
