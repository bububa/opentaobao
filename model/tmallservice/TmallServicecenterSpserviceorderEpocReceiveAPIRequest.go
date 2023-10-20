package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterspserviceorderepocreceiveAPIRequest 电子保单数据接口 API请求
// tmall.servicecenter.spserviceorder.epoc.receive
//
// 电子保单数据回传接口
type TmallservicecenterspserviceorderepocreceiveAPIRequest struct {
	model.Params
	// 电子保单协议号
	_agreementNumber string
	// 电子保单到期时间（包含时分秒）
	_expirationTime string
	// 设备序列号
	_deviceSerialNumber string
	// 产品型号
	_productModel string
	// 服务交易订单号
	_bizOrderId int64
}

// NewTmallservicecenterspserviceorderepocreceiveRequest 初始化TmallservicecenterspserviceorderepocreceiveAPIRequest对象
func NewTmallservicecenterspserviceorderepocreceiveRequest() *TmallservicecenterspserviceorderepocreceiveAPIRequest {
	return &TmallservicecenterspserviceorderepocreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterspserviceorderepocreceiveAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.spserviceorder.epoc.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterspserviceorderepocreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterspserviceorderepocreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgreementNumber is AgreementNumber Setter
// 电子保单协议号
func (r *TmallservicecenterspserviceorderepocreceiveAPIRequest) SetAgreementNumber(_agreementNumber string) error {
	r._agreementNumber = _agreementNumber
	r.Set("agreement_number", _agreementNumber)
	return nil
}

// GetAgreementNumber AgreementNumber Getter
func (r TmallservicecenterspserviceorderepocreceiveAPIRequest) GetAgreementNumber() string {
	return r._agreementNumber
}

// SetExpirationTime is ExpirationTime Setter
// 电子保单到期时间（包含时分秒）
func (r *TmallservicecenterspserviceorderepocreceiveAPIRequest) SetExpirationTime(_expirationTime string) error {
	r._expirationTime = _expirationTime
	r.Set("expiration_time", _expirationTime)
	return nil
}

// GetExpirationTime ExpirationTime Getter
func (r TmallservicecenterspserviceorderepocreceiveAPIRequest) GetExpirationTime() string {
	return r._expirationTime
}

// SetDeviceSerialNumber is DeviceSerialNumber Setter
// 设备序列号
func (r *TmallservicecenterspserviceorderepocreceiveAPIRequest) SetDeviceSerialNumber(_deviceSerialNumber string) error {
	r._deviceSerialNumber = _deviceSerialNumber
	r.Set("device_serial_number", _deviceSerialNumber)
	return nil
}

// GetDeviceSerialNumber DeviceSerialNumber Getter
func (r TmallservicecenterspserviceorderepocreceiveAPIRequest) GetDeviceSerialNumber() string {
	return r._deviceSerialNumber
}

// SetProductModel is ProductModel Setter
// 产品型号
func (r *TmallservicecenterspserviceorderepocreceiveAPIRequest) SetProductModel(_productModel string) error {
	r._productModel = _productModel
	r.Set("product_model", _productModel)
	return nil
}

// GetProductModel ProductModel Getter
func (r TmallservicecenterspserviceorderepocreceiveAPIRequest) GetProductModel() string {
	return r._productModel
}

// SetBizOrderId is BizOrderId Setter
// 服务交易订单号
func (r *TmallservicecenterspserviceorderepocreceiveAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallservicecenterspserviceorderepocreceiveAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
