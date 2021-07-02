package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceMerchantAddAPIRequest 发票中台-同平台授权税号适用商户 API请求
// alibaba.einvoice.merchant.add
//
// 适用于以下场景：
// 业务税号入驻成功后，需要将税号授权给同平台下其他商户，使得其他商户也具备开票能力
type AlibabaEinvoiceMerchantAddAPIRequest struct {
	model.Params
	// 验证码，门店绑定已入驻税号接口返回的taxToken
	_taxToken string
	// 业务方发起新增门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
	_outerId string
	// 业务平台门店ID
	_merchantUserId string
	// 业务平台code, 由阿里发票分配
	_platformCode string
	// 税务登记号
	_payeeRegisterNo string
	// 业务平台门店名称
	_merchantName string
	// 税盘列表
	_deviceIds []string
}

// NewAlibabaEinvoiceMerchantAddRequest 初始化AlibabaEinvoiceMerchantAddAPIRequest对象
func NewAlibabaEinvoiceMerchantAddRequest() *AlibabaEinvoiceMerchantAddAPIRequest {
	return &AlibabaEinvoiceMerchantAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceMerchantAddAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.merchant.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceMerchantAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TaxToken Setter
// 验证码，门店绑定已入驻税号接口返回的taxToken
func (r *AlibabaEinvoiceMerchantAddAPIRequest) SetTaxToken(_taxToken string) error {
	r._taxToken = _taxToken
	r.Set("tax_token", _taxToken)
	return nil
}

// Get TaxToken Getter
func (r AlibabaEinvoiceMerchantAddAPIRequest) GetTaxToken() string {
	return r._taxToken
}

// Set is OuterId Setter
// 业务方发起新增门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
func (r *AlibabaEinvoiceMerchantAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r AlibabaEinvoiceMerchantAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is MerchantUserId Setter
// 业务平台门店ID
func (r *AlibabaEinvoiceMerchantAddAPIRequest) SetMerchantUserId(_merchantUserId string) error {
	r._merchantUserId = _merchantUserId
	r.Set("merchant_user_id", _merchantUserId)
	return nil
}

// Get MerchantUserId Getter
func (r AlibabaEinvoiceMerchantAddAPIRequest) GetMerchantUserId() string {
	return r._merchantUserId
}

// Set is PlatformCode Setter
// 业务平台code, 由阿里发票分配
func (r *AlibabaEinvoiceMerchantAddAPIRequest) SetPlatformCode(_platformCode string) error {
	r._platformCode = _platformCode
	r.Set("platform_code", _platformCode)
	return nil
}

// Get PlatformCode Getter
func (r AlibabaEinvoiceMerchantAddAPIRequest) GetPlatformCode() string {
	return r._platformCode
}

// Set is PayeeRegisterNo Setter
// 税务登记号
func (r *AlibabaEinvoiceMerchantAddAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// Get PayeeRegisterNo Getter
func (r AlibabaEinvoiceMerchantAddAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// Set is MerchantName Setter
// 业务平台门店名称
func (r *AlibabaEinvoiceMerchantAddAPIRequest) SetMerchantName(_merchantName string) error {
	r._merchantName = _merchantName
	r.Set("merchant_name", _merchantName)
	return nil
}

// Get MerchantName Getter
func (r AlibabaEinvoiceMerchantAddAPIRequest) GetMerchantName() string {
	return r._merchantName
}

// Set is DeviceIds Setter
// 税盘列表
func (r *AlibabaEinvoiceMerchantAddAPIRequest) SetDeviceIds(_deviceIds []string) error {
	r._deviceIds = _deviceIds
	r.Set("device_ids", _deviceIds)
	return nil
}

// Get DeviceIds Getter
func (r AlibabaEinvoiceMerchantAddAPIRequest) GetDeviceIds() []string {
	return r._deviceIds
}
