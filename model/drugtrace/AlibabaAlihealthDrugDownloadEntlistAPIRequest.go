package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadEntlistAPIRequest 企业下载列表 API请求
// alibaba.alihealth.drug.download.entlist
//
// 获取企业的下载文件列表
type AlibabaAlihealthDrugDownloadEntlistAPIRequest struct {
	model.Params
	// appKey
	_appKeyN string
}

// NewAlibabaAlihealthDrugDownloadEntlistRequest 初始化AlibabaAlihealthDrugDownloadEntlistAPIRequest对象
func NewAlibabaAlihealthDrugDownloadEntlistRequest() *AlibabaAlihealthDrugDownloadEntlistAPIRequest {
	return &AlibabaAlihealthDrugDownloadEntlistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugDownloadEntlistAPIRequest) Reset() {
	r._appKeyN = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugDownloadEntlistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.entlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugDownloadEntlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugDownloadEntlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppKeyN is AppKeyN Setter
// appKey
func (r *AlibabaAlihealthDrugDownloadEntlistAPIRequest) SetAppKeyN(_appKeyN string) error {
	r._appKeyN = _appKeyN
	r.Set("app_key_n", _appKeyN)
	return nil
}

// GetAppKeyN AppKeyN Getter
func (r AlibabaAlihealthDrugDownloadEntlistAPIRequest) GetAppKeyN() string {
	return r._appKeyN
}

var poolAlibabaAlihealthDrugDownloadEntlistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugDownloadEntlistRequest()
	},
}

// GetAlibabaAlihealthDrugDownloadEntlistRequest 从 sync.Pool 获取 AlibabaAlihealthDrugDownloadEntlistAPIRequest
func GetAlibabaAlihealthDrugDownloadEntlistAPIRequest() *AlibabaAlihealthDrugDownloadEntlistAPIRequest {
	return poolAlibabaAlihealthDrugDownloadEntlistAPIRequest.Get().(*AlibabaAlihealthDrugDownloadEntlistAPIRequest)
}

// ReleaseAlibabaAlihealthDrugDownloadEntlistAPIRequest 将 AlibabaAlihealthDrugDownloadEntlistAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugDownloadEntlistAPIRequest(v *AlibabaAlihealthDrugDownloadEntlistAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugDownloadEntlistAPIRequest.Put(v)
}
