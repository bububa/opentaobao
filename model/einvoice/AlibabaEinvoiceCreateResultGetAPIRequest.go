package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceCreateResultGetAPIRequest ERP开票结果获取 API请求
// alibaba.einvoice.create.result.get
//
// ERP开票结果获取
type AlibabaEinvoiceCreateResultGetAPIRequest struct {
	model.Params
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 电商平台代码。淘宝：taobao，天猫：tmall
	_platformCode string
	// 电商平台对应的订单号
	_platformTid string
	// 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
	_serialNo string
	// 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
	_outShopName string
}

// NewAlibabaEinvoiceCreateResultGetRequest 初始化AlibabaEinvoiceCreateResultGetAPIRequest对象
func NewAlibabaEinvoiceCreateResultGetRequest() *AlibabaEinvoiceCreateResultGetAPIRequest {
	return &AlibabaEinvoiceCreateResultGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceCreateResultGetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.create.result.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceCreateResultGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceCreateResultGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceCreateResultGetAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceCreateResultGetAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetPlatformCode is PlatformCode Setter
// 电商平台代码。淘宝：taobao，天猫：tmall
func (r *AlibabaEinvoiceCreateResultGetAPIRequest) SetPlatformCode(_platformCode string) error {
	r._platformCode = _platformCode
	r.Set("platform_code", _platformCode)
	return nil
}

// GetPlatformCode PlatformCode Getter
func (r AlibabaEinvoiceCreateResultGetAPIRequest) GetPlatformCode() string {
	return r._platformCode
}

// SetPlatformTid is PlatformTid Setter
// 电商平台对应的订单号
func (r *AlibabaEinvoiceCreateResultGetAPIRequest) SetPlatformTid(_platformTid string) error {
	r._platformTid = _platformTid
	r.Set("platform_tid", _platformTid)
	return nil
}

// GetPlatformTid PlatformTid Getter
func (r AlibabaEinvoiceCreateResultGetAPIRequest) GetPlatformTid() string {
	return r._platformTid
}

// SetSerialNo is SerialNo Setter
// 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
func (r *AlibabaEinvoiceCreateResultGetAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaEinvoiceCreateResultGetAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetOutShopName is OutShopName Setter
// 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
func (r *AlibabaEinvoiceCreateResultGetAPIRequest) SetOutShopName(_outShopName string) error {
	r._outShopName = _outShopName
	r.Set("out_shop_name", _outShopName)
	return nil
}

// GetOutShopName OutShopName Getter
func (r AlibabaEinvoiceCreateResultGetAPIRequest) GetOutShopName() string {
	return r._outShopName
}
