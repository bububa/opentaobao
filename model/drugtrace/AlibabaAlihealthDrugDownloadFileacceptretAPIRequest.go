package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadFileacceptretAPIRequest 企业上传回执 API请求
// alibaba.alihealth.drug.download.fileacceptret
//
// 拿到企业下载回执，将企业已下载的和未下载成功的条目都相应的改变状态
type AlibabaAlihealthDrugDownloadFileacceptretAPIRequest struct {
	model.Params
	// appKey
	_appKeyN string
	// fileResultJson
	_fileResultJson string
}

// NewAlibabaAlihealthDrugDownloadFileacceptretRequest 初始化AlibabaAlihealthDrugDownloadFileacceptretAPIRequest对象
func NewAlibabaAlihealthDrugDownloadFileacceptretRequest() *AlibabaAlihealthDrugDownloadFileacceptretAPIRequest {
	return &AlibabaAlihealthDrugDownloadFileacceptretAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugDownloadFileacceptretAPIRequest) Reset() {
	r._appKeyN = ""
	r._fileResultJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugDownloadFileacceptretAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.fileacceptret"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugDownloadFileacceptretAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugDownloadFileacceptretAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppKeyN is AppKeyN Setter
// appKey
func (r *AlibabaAlihealthDrugDownloadFileacceptretAPIRequest) SetAppKeyN(_appKeyN string) error {
	r._appKeyN = _appKeyN
	r.Set("app_key_n", _appKeyN)
	return nil
}

// GetAppKeyN AppKeyN Getter
func (r AlibabaAlihealthDrugDownloadFileacceptretAPIRequest) GetAppKeyN() string {
	return r._appKeyN
}

// SetFileResultJson is FileResultJson Setter
// fileResultJson
func (r *AlibabaAlihealthDrugDownloadFileacceptretAPIRequest) SetFileResultJson(_fileResultJson string) error {
	r._fileResultJson = _fileResultJson
	r.Set("file_result_json", _fileResultJson)
	return nil
}

// GetFileResultJson FileResultJson Getter
func (r AlibabaAlihealthDrugDownloadFileacceptretAPIRequest) GetFileResultJson() string {
	return r._fileResultJson
}

var poolAlibabaAlihealthDrugDownloadFileacceptretAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugDownloadFileacceptretRequest()
	},
}

// GetAlibabaAlihealthDrugDownloadFileacceptretRequest 从 sync.Pool 获取 AlibabaAlihealthDrugDownloadFileacceptretAPIRequest
func GetAlibabaAlihealthDrugDownloadFileacceptretAPIRequest() *AlibabaAlihealthDrugDownloadFileacceptretAPIRequest {
	return poolAlibabaAlihealthDrugDownloadFileacceptretAPIRequest.Get().(*AlibabaAlihealthDrugDownloadFileacceptretAPIRequest)
}

// ReleaseAlibabaAlihealthDrugDownloadFileacceptretAPIRequest 将 AlibabaAlihealthDrugDownloadFileacceptretAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugDownloadFileacceptretAPIRequest(v *AlibabaAlihealthDrugDownloadFileacceptretAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugDownloadFileacceptretAPIRequest.Put(v)
}
