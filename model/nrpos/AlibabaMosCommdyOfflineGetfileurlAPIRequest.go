package nrpos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosCommdyOfflineGetfileurlAPIRequest 去前置机pos商品离线文件下载地址查询接口 API请求
// alibaba.mos.commdy.offline.getfileurl
//
// 去前置机-pos查询离线文件下载地址接口
type AlibabaMosCommdyOfflineGetfileurlAPIRequest struct {
	model.Params
	// 离线文件名称
	_fileKeys []string
}

// NewAlibabaMosCommdyOfflineGetfileurlRequest 初始化AlibabaMosCommdyOfflineGetfileurlAPIRequest对象
func NewAlibabaMosCommdyOfflineGetfileurlRequest() *AlibabaMosCommdyOfflineGetfileurlAPIRequest {
	return &AlibabaMosCommdyOfflineGetfileurlAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosCommdyOfflineGetfileurlAPIRequest) Reset() {
	r._fileKeys = r._fileKeys[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosCommdyOfflineGetfileurlAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.commdy.offline.getfileurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosCommdyOfflineGetfileurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosCommdyOfflineGetfileurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileKeys is FileKeys Setter
// 离线文件名称
func (r *AlibabaMosCommdyOfflineGetfileurlAPIRequest) SetFileKeys(_fileKeys []string) error {
	r._fileKeys = _fileKeys
	r.Set("file_keys", _fileKeys)
	return nil
}

// GetFileKeys FileKeys Getter
func (r AlibabaMosCommdyOfflineGetfileurlAPIRequest) GetFileKeys() []string {
	return r._fileKeys
}

var poolAlibabaMosCommdyOfflineGetfileurlAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosCommdyOfflineGetfileurlRequest()
	},
}

// GetAlibabaMosCommdyOfflineGetfileurlRequest 从 sync.Pool 获取 AlibabaMosCommdyOfflineGetfileurlAPIRequest
func GetAlibabaMosCommdyOfflineGetfileurlAPIRequest() *AlibabaMosCommdyOfflineGetfileurlAPIRequest {
	return poolAlibabaMosCommdyOfflineGetfileurlAPIRequest.Get().(*AlibabaMosCommdyOfflineGetfileurlAPIRequest)
}

// ReleaseAlibabaMosCommdyOfflineGetfileurlAPIRequest 将 AlibabaMosCommdyOfflineGetfileurlAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosCommdyOfflineGetfileurlAPIRequest(v *AlibabaMosCommdyOfflineGetfileurlAPIRequest) {
	v.Reset()
	poolAlibabaMosCommdyOfflineGetfileurlAPIRequest.Put(v)
}
