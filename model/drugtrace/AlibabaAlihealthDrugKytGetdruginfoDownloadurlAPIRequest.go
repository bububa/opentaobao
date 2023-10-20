package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest 药品全量数据下载 API请求
// alibaba.alihealth.drug.kyt.getdruginfo.downloadurl
//
// 药品全量数据下载
type AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest struct {
	model.Params
	// 调用接口的企业ID
	_refEntId string
}

// NewAlibabaalihealthdrugkytgetdruginfodownloadurlRequest 初始化AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest对象
func NewAlibabaalihealthdrugkytgetdruginfodownloadurlRequest() *AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest {
	return &AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getdruginfo.downloadurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用接口的企业ID
func (r *AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest) GetRefEntId() string {
	return r._refEntId
}
