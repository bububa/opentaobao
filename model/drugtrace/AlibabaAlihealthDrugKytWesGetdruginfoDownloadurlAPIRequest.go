package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest 药品全量数据下载 API请求
// alibaba.alihealth.drug.kyt.wes.getdruginfo.downloadurl
//
// 药品全量数据下载
type AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest struct {
	model.Params
	// 调用接口的企业ID
	_refEntId string
	// 服务时间
	_licenseToken string
}

// NewAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest 初始化AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest对象
func NewAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest() *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest {
	return &AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.getdruginfo.downloadurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用接口的企业ID
func (r *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}
