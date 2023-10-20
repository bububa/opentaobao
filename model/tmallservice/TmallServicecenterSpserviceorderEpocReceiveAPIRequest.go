package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderEpocReceiveAPIRequest 电子保单数据接口 API请求
// tmall.servicecenter.spserviceorder.epoc.receive
//
// 电子保单数据回传接口
type TmallServicecenterSpserviceorderEpocReceiveAPIRequest struct {
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

// NewTmallServicecenterSpserviceorderEpocReceiveRequest 初始化TmallServicecenterSpserviceorderEpocReceiveAPIRequest对象
func NewTmallServicecenterSpserviceorderEpocReceiveRequest() *TmallServicecenterSpserviceorderEpocReceiveAPIRequest {
	return &TmallServicecenterSpserviceorderEpocReceiveAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterSpserviceorderEpocReceiveAPIRequest) Reset() {
	r._agreementNumber = ""
	r._expirationTime = ""
	r._deviceSerialNumber = ""
	r._productModel = ""
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterSpserviceorderEpocReceiveAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.spserviceorder.epoc.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterSpserviceorderEpocReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterSpserviceorderEpocReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgreementNumber is AgreementNumber Setter
// 电子保单协议号
func (r *TmallServicecenterSpserviceorderEpocReceiveAPIRequest) SetAgreementNumber(_agreementNumber string) error {
	r._agreementNumber = _agreementNumber
	r.Set("agreement_number", _agreementNumber)
	return nil
}

// GetAgreementNumber AgreementNumber Getter
func (r TmallServicecenterSpserviceorderEpocReceiveAPIRequest) GetAgreementNumber() string {
	return r._agreementNumber
}

// SetExpirationTime is ExpirationTime Setter
// 电子保单到期时间（包含时分秒）
func (r *TmallServicecenterSpserviceorderEpocReceiveAPIRequest) SetExpirationTime(_expirationTime string) error {
	r._expirationTime = _expirationTime
	r.Set("expiration_time", _expirationTime)
	return nil
}

// GetExpirationTime ExpirationTime Getter
func (r TmallServicecenterSpserviceorderEpocReceiveAPIRequest) GetExpirationTime() string {
	return r._expirationTime
}

// SetDeviceSerialNumber is DeviceSerialNumber Setter
// 设备序列号
func (r *TmallServicecenterSpserviceorderEpocReceiveAPIRequest) SetDeviceSerialNumber(_deviceSerialNumber string) error {
	r._deviceSerialNumber = _deviceSerialNumber
	r.Set("device_serial_number", _deviceSerialNumber)
	return nil
}

// GetDeviceSerialNumber DeviceSerialNumber Getter
func (r TmallServicecenterSpserviceorderEpocReceiveAPIRequest) GetDeviceSerialNumber() string {
	return r._deviceSerialNumber
}

// SetProductModel is ProductModel Setter
// 产品型号
func (r *TmallServicecenterSpserviceorderEpocReceiveAPIRequest) SetProductModel(_productModel string) error {
	r._productModel = _productModel
	r.Set("product_model", _productModel)
	return nil
}

// GetProductModel ProductModel Getter
func (r TmallServicecenterSpserviceorderEpocReceiveAPIRequest) GetProductModel() string {
	return r._productModel
}

// SetBizOrderId is BizOrderId Setter
// 服务交易订单号
func (r *TmallServicecenterSpserviceorderEpocReceiveAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallServicecenterSpserviceorderEpocReceiveAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolTmallServicecenterSpserviceorderEpocReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterSpserviceorderEpocReceiveRequest()
	},
}

// GetTmallServicecenterSpserviceorderEpocReceiveRequest 从 sync.Pool 获取 TmallServicecenterSpserviceorderEpocReceiveAPIRequest
func GetTmallServicecenterSpserviceorderEpocReceiveAPIRequest() *TmallServicecenterSpserviceorderEpocReceiveAPIRequest {
	return poolTmallServicecenterSpserviceorderEpocReceiveAPIRequest.Get().(*TmallServicecenterSpserviceorderEpocReceiveAPIRequest)
}

// ReleaseTmallServicecenterSpserviceorderEpocReceiveAPIRequest 将 TmallServicecenterSpserviceorderEpocReceiveAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterSpserviceorderEpocReceiveAPIRequest(v *TmallServicecenterSpserviceorderEpocReceiveAPIRequest) {
	v.Reset()
	poolTmallServicecenterSpserviceorderEpocReceiveAPIRequest.Put(v)
}
