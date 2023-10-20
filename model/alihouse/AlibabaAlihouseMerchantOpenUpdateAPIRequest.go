package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousemerchantopenupdateAPIRequest 非融合店进件升级成融合店 API请求
// alibaba.alihouse.merchant.open.update
//
// 非融合店进件升级成融合店
type AlibabaalihousemerchantopenupdateAPIRequest struct {
	model.Params
	// 法人证件有效期时间
	_legalCertExpireTime string
	// 经营主体ID
	_merchantOpenId int64
	// 法人证件有效期类型
	_legalCertStatus int64
}

// NewAlibabaalihousemerchantopenupdateRequest 初始化AlibabaalihousemerchantopenupdateAPIRequest对象
func NewAlibabaalihousemerchantopenupdateRequest() *AlibabaalihousemerchantopenupdateAPIRequest {
	return &AlibabaalihousemerchantopenupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousemerchantopenupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.merchant.open.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousemerchantopenupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousemerchantopenupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLegalCertExpireTime is LegalCertExpireTime Setter
// 法人证件有效期时间
func (r *AlibabaalihousemerchantopenupdateAPIRequest) SetLegalCertExpireTime(_legalCertExpireTime string) error {
	r._legalCertExpireTime = _legalCertExpireTime
	r.Set("legal_cert_expire_time", _legalCertExpireTime)
	return nil
}

// GetLegalCertExpireTime LegalCertExpireTime Getter
func (r AlibabaalihousemerchantopenupdateAPIRequest) GetLegalCertExpireTime() string {
	return r._legalCertExpireTime
}

// SetMerchantOpenId is MerchantOpenId Setter
// 经营主体ID
func (r *AlibabaalihousemerchantopenupdateAPIRequest) SetMerchantOpenId(_merchantOpenId int64) error {
	r._merchantOpenId = _merchantOpenId
	r.Set("merchant_open_id", _merchantOpenId)
	return nil
}

// GetMerchantOpenId MerchantOpenId Getter
func (r AlibabaalihousemerchantopenupdateAPIRequest) GetMerchantOpenId() int64 {
	return r._merchantOpenId
}

// SetLegalCertStatus is LegalCertStatus Setter
// 法人证件有效期类型
func (r *AlibabaalihousemerchantopenupdateAPIRequest) SetLegalCertStatus(_legalCertStatus int64) error {
	r._legalCertStatus = _legalCertStatus
	r.Set("legal_cert_status", _legalCertStatus)
	return nil
}

// GetLegalCertStatus LegalCertStatus Getter
func (r AlibabaalihousemerchantopenupdateAPIRequest) GetLegalCertStatus() int64 {
	return r._legalCertStatus
}
