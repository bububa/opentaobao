package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicemerchantaddAPIRequest 发票中台-同平台授权税号适用商户 API请求
// alibaba.einvoice.merchant.add
//
// 适用于以下场景：
// 业务税号入驻成功后，需要将税号授权给同平台下其他商户，使得其他商户也具备开票能力
type AlibabaeinvoicemerchantaddAPIRequest struct {
	model.Params
	// 税盘列表
	_deviceIds []string
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
}

// NewAlibabaeinvoicemerchantaddRequest 初始化AlibabaeinvoicemerchantaddAPIRequest对象
func NewAlibabaeinvoicemerchantaddRequest() *AlibabaeinvoicemerchantaddAPIRequest {
	return &AlibabaeinvoicemerchantaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicemerchantaddAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.merchant.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicemerchantaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicemerchantaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceIds is DeviceIds Setter
// 税盘列表
func (r *AlibabaeinvoicemerchantaddAPIRequest) SetDeviceIds(_deviceIds []string) error {
	r._deviceIds = _deviceIds
	r.Set("device_ids", _deviceIds)
	return nil
}

// GetDeviceIds DeviceIds Getter
func (r AlibabaeinvoicemerchantaddAPIRequest) GetDeviceIds() []string {
	return r._deviceIds
}

// SetTaxToken is TaxToken Setter
// 验证码，门店绑定已入驻税号接口返回的taxToken
func (r *AlibabaeinvoicemerchantaddAPIRequest) SetTaxToken(_taxToken string) error {
	r._taxToken = _taxToken
	r.Set("tax_token", _taxToken)
	return nil
}

// GetTaxToken TaxToken Getter
func (r AlibabaeinvoicemerchantaddAPIRequest) GetTaxToken() string {
	return r._taxToken
}

// SetOuterId is OuterId Setter
// 业务方发起新增门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
func (r *AlibabaeinvoicemerchantaddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaeinvoicemerchantaddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetMerchantUserId is MerchantUserId Setter
// 业务平台门店ID
func (r *AlibabaeinvoicemerchantaddAPIRequest) SetMerchantUserId(_merchantUserId string) error {
	r._merchantUserId = _merchantUserId
	r.Set("merchant_user_id", _merchantUserId)
	return nil
}

// GetMerchantUserId MerchantUserId Getter
func (r AlibabaeinvoicemerchantaddAPIRequest) GetMerchantUserId() string {
	return r._merchantUserId
}

// SetPlatformCode is PlatformCode Setter
// 业务平台code, 由阿里发票分配
func (r *AlibabaeinvoicemerchantaddAPIRequest) SetPlatformCode(_platformCode string) error {
	r._platformCode = _platformCode
	r.Set("platform_code", _platformCode)
	return nil
}

// GetPlatformCode PlatformCode Getter
func (r AlibabaeinvoicemerchantaddAPIRequest) GetPlatformCode() string {
	return r._platformCode
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 税务登记号
func (r *AlibabaeinvoicemerchantaddAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoicemerchantaddAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetMerchantName is MerchantName Setter
// 业务平台门店名称
func (r *AlibabaeinvoicemerchantaddAPIRequest) SetMerchantName(_merchantName string) error {
	r._merchantName = _merchantName
	r.Set("merchant_name", _merchantName)
	return nil
}

// GetMerchantName MerchantName Getter
func (r AlibabaeinvoicemerchantaddAPIRequest) GetMerchantName() string {
	return r._merchantName
}
