package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdownloadfileacceptretAPIRequest 企业上传回执 API请求
// alibaba.alihealth.drug.download.fileacceptret
//
// 拿到企业下载回执，将企业已下载的和未下载成功的条目都相应的改变状态
type AlibabaalihealthdrugdownloadfileacceptretAPIRequest struct {
	model.Params
	// appKey
	_appKeyN string
	// fileResultJson
	_fileResultJson string
}

// NewAlibabaalihealthdrugdownloadfileacceptretRequest 初始化AlibabaalihealthdrugdownloadfileacceptretAPIRequest对象
func NewAlibabaalihealthdrugdownloadfileacceptretRequest() *AlibabaalihealthdrugdownloadfileacceptretAPIRequest {
	return &AlibabaalihealthdrugdownloadfileacceptretAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugdownloadfileacceptretAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.fileacceptret"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugdownloadfileacceptretAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugdownloadfileacceptretAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppKeyN is AppKeyN Setter
// appKey
func (r *AlibabaalihealthdrugdownloadfileacceptretAPIRequest) SetAppKeyN(_appKeyN string) error {
	r._appKeyN = _appKeyN
	r.Set("app_key_n", _appKeyN)
	return nil
}

// GetAppKeyN AppKeyN Getter
func (r AlibabaalihealthdrugdownloadfileacceptretAPIRequest) GetAppKeyN() string {
	return r._appKeyN
}

// SetFileResultJson is FileResultJson Setter
// fileResultJson
func (r *AlibabaalihealthdrugdownloadfileacceptretAPIRequest) SetFileResultJson(_fileResultJson string) error {
	r._fileResultJson = _fileResultJson
	r.Set("file_result_json", _fileResultJson)
	return nil
}

// GetFileResultJson FileResultJson Getter
func (r AlibabaalihealthdrugdownloadfileacceptretAPIRequest) GetFileResultJson() string {
	return r._fileResultJson
}
