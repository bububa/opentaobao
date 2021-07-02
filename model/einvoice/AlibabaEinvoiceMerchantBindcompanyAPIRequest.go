package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceMerchantBindcompanyAPIRequest 发票中台-跨平台绑定已入驻税号与商户 API请求
// alibaba.einvoice.merchant.bindcompany
//
// 税号在阿里发票平台入驻成功后，允许业务方通过本接口跨业务平台绑定入驻税号和业务平台商户，绑定成功后该商户可以使用该税号的盘进行开票。绑定成功后，可以使用同平台授权、取消授权税号适用商户接口来变更税号和商户关系。
type AlibabaEinvoiceMerchantBindcompanyAPIRequest struct {
	model.Params
	// 业务方发起首次绑定门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
	_outerId string
	// 业务平台商户ID
	_merchantUserId string
	// 激活码
	_activationCode string
	// 业务平台code, 由阿里发票分配
	_platformCode string
	// 税务登记号
	_payeeRegisterNo string
	// 业务平台门店名称
	_merchantName string
	// 税号已入驻的原业务平台code
	_sourcePlatformCode string
}

// NewAlibabaEinvoiceMerchantBindcompanyRequest 初始化AlibabaEinvoiceMerchantBindcompanyAPIRequest对象
func NewAlibabaEinvoiceMerchantBindcompanyRequest() *AlibabaEinvoiceMerchantBindcompanyAPIRequest {
	return &AlibabaEinvoiceMerchantBindcompanyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceMerchantBindcompanyAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.merchant.bindcompany"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceMerchantBindcompanyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterId is OuterId Setter
// 业务方发起首次绑定门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
func (r *AlibabaEinvoiceMerchantBindcompanyAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaEinvoiceMerchantBindcompanyAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetMerchantUserId is MerchantUserId Setter
// 业务平台商户ID
func (r *AlibabaEinvoiceMerchantBindcompanyAPIRequest) SetMerchantUserId(_merchantUserId string) error {
	r._merchantUserId = _merchantUserId
	r.Set("merchant_user_id", _merchantUserId)
	return nil
}

// GetMerchantUserId MerchantUserId Getter
func (r AlibabaEinvoiceMerchantBindcompanyAPIRequest) GetMerchantUserId() string {
	return r._merchantUserId
}

// SetActivationCode is ActivationCode Setter
// 激活码
func (r *AlibabaEinvoiceMerchantBindcompanyAPIRequest) SetActivationCode(_activationCode string) error {
	r._activationCode = _activationCode
	r.Set("activation_code", _activationCode)
	return nil
}

// GetActivationCode ActivationCode Getter
func (r AlibabaEinvoiceMerchantBindcompanyAPIRequest) GetActivationCode() string {
	return r._activationCode
}

// SetPlatformCode is PlatformCode Setter
// 业务平台code, 由阿里发票分配
func (r *AlibabaEinvoiceMerchantBindcompanyAPIRequest) SetPlatformCode(_platformCode string) error {
	r._platformCode = _platformCode
	r.Set("platform_code", _platformCode)
	return nil
}

// GetPlatformCode PlatformCode Getter
func (r AlibabaEinvoiceMerchantBindcompanyAPIRequest) GetPlatformCode() string {
	return r._platformCode
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 税务登记号
func (r *AlibabaEinvoiceMerchantBindcompanyAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceMerchantBindcompanyAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetMerchantName is MerchantName Setter
// 业务平台门店名称
func (r *AlibabaEinvoiceMerchantBindcompanyAPIRequest) SetMerchantName(_merchantName string) error {
	r._merchantName = _merchantName
	r.Set("merchant_name", _merchantName)
	return nil
}

// GetMerchantName MerchantName Getter
func (r AlibabaEinvoiceMerchantBindcompanyAPIRequest) GetMerchantName() string {
	return r._merchantName
}

// SetSourcePlatformCode is SourcePlatformCode Setter
// 税号已入驻的原业务平台code
func (r *AlibabaEinvoiceMerchantBindcompanyAPIRequest) SetSourcePlatformCode(_sourcePlatformCode string) error {
	r._sourcePlatformCode = _sourcePlatformCode
	r.Set("source_platform_code", _sourcePlatformCode)
	return nil
}

// GetSourcePlatformCode SourcePlatformCode Getter
func (r AlibabaEinvoiceMerchantBindcompanyAPIRequest) GetSourcePlatformCode() string {
	return r._sourcePlatformCode
}
