package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest 获取重点追溯品种明细下载URL API请求
// alibaba.alihealth.drugtrace.top.yljg.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
type AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest struct {
	model.Params
	// 调用接口的企业ID
	_refEntId string
}

// NewAlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlRequest 初始化AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest对象
func NewAlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlRequest() *AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest {
	return &AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.getkeyflagdruginfo.downloadurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用接口的企业ID
func (r *AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest) GetRefEntId() string {
	return r._refEntId
}
