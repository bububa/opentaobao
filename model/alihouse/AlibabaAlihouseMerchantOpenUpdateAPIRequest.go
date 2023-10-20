package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseMerchantOpenUpdateAPIRequest 非融合店进件升级成融合店 API请求
// alibaba.alihouse.merchant.open.update
//
// 非融合店进件升级成融合店
type AlibabaAlihouseMerchantOpenUpdateAPIRequest struct {
	model.Params
	// 法人证件有效期时间
	_legalCertExpireTime string
	// 经营主体ID
	_merchantOpenId int64
	// 法人证件有效期类型
	_legalCertStatus int64
}

// NewAlibabaAlihouseMerchantOpenUpdateRequest 初始化AlibabaAlihouseMerchantOpenUpdateAPIRequest对象
func NewAlibabaAlihouseMerchantOpenUpdateRequest() *AlibabaAlihouseMerchantOpenUpdateAPIRequest {
	return &AlibabaAlihouseMerchantOpenUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseMerchantOpenUpdateAPIRequest) Reset() {
	r._legalCertExpireTime = ""
	r._merchantOpenId = 0
	r._legalCertStatus = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseMerchantOpenUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.merchant.open.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseMerchantOpenUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseMerchantOpenUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLegalCertExpireTime is LegalCertExpireTime Setter
// 法人证件有效期时间
func (r *AlibabaAlihouseMerchantOpenUpdateAPIRequest) SetLegalCertExpireTime(_legalCertExpireTime string) error {
	r._legalCertExpireTime = _legalCertExpireTime
	r.Set("legal_cert_expire_time", _legalCertExpireTime)
	return nil
}

// GetLegalCertExpireTime LegalCertExpireTime Getter
func (r AlibabaAlihouseMerchantOpenUpdateAPIRequest) GetLegalCertExpireTime() string {
	return r._legalCertExpireTime
}

// SetMerchantOpenId is MerchantOpenId Setter
// 经营主体ID
func (r *AlibabaAlihouseMerchantOpenUpdateAPIRequest) SetMerchantOpenId(_merchantOpenId int64) error {
	r._merchantOpenId = _merchantOpenId
	r.Set("merchant_open_id", _merchantOpenId)
	return nil
}

// GetMerchantOpenId MerchantOpenId Getter
func (r AlibabaAlihouseMerchantOpenUpdateAPIRequest) GetMerchantOpenId() int64 {
	return r._merchantOpenId
}

// SetLegalCertStatus is LegalCertStatus Setter
// 法人证件有效期类型
func (r *AlibabaAlihouseMerchantOpenUpdateAPIRequest) SetLegalCertStatus(_legalCertStatus int64) error {
	r._legalCertStatus = _legalCertStatus
	r.Set("legal_cert_status", _legalCertStatus)
	return nil
}

// GetLegalCertStatus LegalCertStatus Getter
func (r AlibabaAlihouseMerchantOpenUpdateAPIRequest) GetLegalCertStatus() int64 {
	return r._legalCertStatus
}

var poolAlibabaAlihouseMerchantOpenUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseMerchantOpenUpdateRequest()
	},
}

// GetAlibabaAlihouseMerchantOpenUpdateRequest 从 sync.Pool 获取 AlibabaAlihouseMerchantOpenUpdateAPIRequest
func GetAlibabaAlihouseMerchantOpenUpdateAPIRequest() *AlibabaAlihouseMerchantOpenUpdateAPIRequest {
	return poolAlibabaAlihouseMerchantOpenUpdateAPIRequest.Get().(*AlibabaAlihouseMerchantOpenUpdateAPIRequest)
}

// ReleaseAlibabaAlihouseMerchantOpenUpdateAPIRequest 将 AlibabaAlihouseMerchantOpenUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseMerchantOpenUpdateAPIRequest(v *AlibabaAlihouseMerchantOpenUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseMerchantOpenUpdateAPIRequest.Put(v)
}
