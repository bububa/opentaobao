package nrpos

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosCommdyOfflineGetfileurlAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.commdy.offline.getfileurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosCommdyOfflineGetfileurlAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FileKeys Setter
// 离线文件名称
func (r *AlibabaMosCommdyOfflineGetfileurlAPIRequest) SetFileKeys(_fileKeys []string) error {
	r._fileKeys = _fileKeys
	r.Set("file_keys", _fileKeys)
	return nil
}

// Get FileKeys Getter
func (r AlibabaMosCommdyOfflineGetfileurlAPIRequest) GetFileKeys() []string {
	return r._fileKeys
}
