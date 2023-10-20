package nrpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoscommdyofflinegetfileurlAPIRequest 去前置机pos商品离线文件下载地址查询接口 API请求
// alibaba.mos.commdy.offline.getfileurl
//
// 去前置机-pos查询离线文件下载地址接口
type AlibabamoscommdyofflinegetfileurlAPIRequest struct {
	model.Params
	// 离线文件名称
	_fileKeys []string
}

// NewAlibabamoscommdyofflinegetfileurlRequest 初始化AlibabamoscommdyofflinegetfileurlAPIRequest对象
func NewAlibabamoscommdyofflinegetfileurlRequest() *AlibabamoscommdyofflinegetfileurlAPIRequest {
	return &AlibabamoscommdyofflinegetfileurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoscommdyofflinegetfileurlAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.commdy.offline.getfileurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoscommdyofflinegetfileurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoscommdyofflinegetfileurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileKeys is FileKeys Setter
// 离线文件名称
func (r *AlibabamoscommdyofflinegetfileurlAPIRequest) SetFileKeys(_fileKeys []string) error {
	r._fileKeys = _fileKeys
	r.Set("file_keys", _fileKeys)
	return nil
}

// GetFileKeys FileKeys Getter
func (r AlibabamoscommdyofflinegetfileurlAPIRequest) GetFileKeys() []string {
	return r._fileKeys
}
