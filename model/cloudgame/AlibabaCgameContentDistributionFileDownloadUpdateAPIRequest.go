package cloudgame

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest 文件下载回调 API请求
// alibaba.cgame.content.distribution.file.download.update
//
// 提供内容服务器更新文件下载状态的能力
type AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_reqParam *FileDownloadCallbackRequest
}

// NewAlibabaCgameContentDistributionFileDownloadUpdateRequest 初始化AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest对象
func NewAlibabaCgameContentDistributionFileDownloadUpdateRequest() *AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest {
	return &AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest) Reset() {
	r._reqParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.content.distribution.file.download.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqParam is ReqParam Setter
// 请求参数
func (r *AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest) SetReqParam(_reqParam *FileDownloadCallbackRequest) error {
	r._reqParam = _reqParam
	r.Set("req_param", _reqParam)
	return nil
}

// GetReqParam ReqParam Getter
func (r AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest) GetReqParam() *FileDownloadCallbackRequest {
	return r._reqParam
}

var poolAlibabaCgameContentDistributionFileDownloadUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCgameContentDistributionFileDownloadUpdateRequest()
	},
}

// GetAlibabaCgameContentDistributionFileDownloadUpdateRequest 从 sync.Pool 获取 AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest
func GetAlibabaCgameContentDistributionFileDownloadUpdateAPIRequest() *AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest {
	return poolAlibabaCgameContentDistributionFileDownloadUpdateAPIRequest.Get().(*AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest)
}

// ReleaseAlibabaCgameContentDistributionFileDownloadUpdateAPIRequest 将 AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaCgameContentDistributionFileDownloadUpdateAPIRequest(v *AlibabaCgameContentDistributionFileDownloadUpdateAPIRequest) {
	v.Reset()
	poolAlibabaCgameContentDistributionFileDownloadUpdateAPIRequest.Put(v)
}
