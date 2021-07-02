package drugtrace

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugDownloadEntlistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.entlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugDownloadEntlistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppKeyN Setter
// appKey
func (r *AlibabaAlihealthDrugDownloadEntlistAPIRequest) SetAppKeyN(_appKeyN string) error {
	r._appKeyN = _appKeyN
	r.Set("app_key_n", _appKeyN)
	return nil
}

// Get AppKeyN Getter
func (r AlibabaAlihealthDrugDownloadEntlistAPIRequest) GetAppKeyN() string {
	return r._appKeyN
}
