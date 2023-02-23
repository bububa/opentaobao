package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest 获取product.xml的下载链接 API请求
// alibaba.alihealth.drug.download.getproductxml
//
// 阿里健康-追溯码-D2D获得药器信息下载地址，方便生产线操作
type AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest struct {
	model.Params
	// appKey
	_appKeyN string
	// 企业名称
	_entName string
}

// NewAlibabaAlihealthDrugDownloadGetproductxmlRequest 初始化AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest对象
func NewAlibabaAlihealthDrugDownloadGetproductxmlRequest() *AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest {
	return &AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.getproductxml"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppKeyN is AppKeyN Setter
// appKey
func (r *AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest) SetAppKeyN(_appKeyN string) error {
	r._appKeyN = _appKeyN
	r.Set("app_key_n", _appKeyN)
	return nil
}

// GetAppKeyN AppKeyN Getter
func (r AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest) GetAppKeyN() string {
	return r._appKeyN
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugDownloadGetproductxmlAPIRequest) GetEntName() string {
	return r._entName
}
