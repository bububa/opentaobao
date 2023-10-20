package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotvpaypromotioninfogetAPIRequest tv支付查询消费抽奖配置 API请求
// taobao.tvpay.promotion.info.get
//
// 查询消费抽奖配置
type TaobaotvpaypromotioninfogetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 淘系订单号
	_extOrderId string
	// 标题
	_subject string
	// 商品id
	_subjectId string
	// 是否淘系
	_isTao bool
}

// NewTaobaotvpaypromotioninfogetRequest 初始化TaobaotvpaypromotioninfogetAPIRequest对象
func NewTaobaotvpaypromotioninfogetRequest() *TaobaotvpaypromotioninfogetAPIRequest {
	return &TaobaotvpaypromotioninfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotvpaypromotioninfogetAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.promotion.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotvpaypromotioninfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotvpaypromotioninfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaotvpaypromotioninfogetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaotvpaypromotioninfogetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaotvpaypromotioninfogetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaotvpaypromotioninfogetAPIRequest) GetFrom() string {
	return r._from
}

// SetExtOrderId is ExtOrderId Setter
// 淘系订单号
func (r *TaobaotvpaypromotioninfogetAPIRequest) SetExtOrderId(_extOrderId string) error {
	r._extOrderId = _extOrderId
	r.Set("ext_order_id", _extOrderId)
	return nil
}

// GetExtOrderId ExtOrderId Getter
func (r TaobaotvpaypromotioninfogetAPIRequest) GetExtOrderId() string {
	return r._extOrderId
}

// SetSubject is Subject Setter
// 标题
func (r *TaobaotvpaypromotioninfogetAPIRequest) SetSubject(_subject string) error {
	r._subject = _subject
	r.Set("subject", _subject)
	return nil
}

// GetSubject Subject Getter
func (r TaobaotvpaypromotioninfogetAPIRequest) GetSubject() string {
	return r._subject
}

// SetSubjectId is SubjectId Setter
// 商品id
func (r *TaobaotvpaypromotioninfogetAPIRequest) SetSubjectId(_subjectId string) error {
	r._subjectId = _subjectId
	r.Set("subject_id", _subjectId)
	return nil
}

// GetSubjectId SubjectId Getter
func (r TaobaotvpaypromotioninfogetAPIRequest) GetSubjectId() string {
	return r._subjectId
}

// SetIsTao is IsTao Setter
// 是否淘系
func (r *TaobaotvpaypromotioninfogetAPIRequest) SetIsTao(_isTao bool) error {
	r._isTao = _isTao
	r.Set("is_tao", _isTao)
	return nil
}

// GetIsTao IsTao Getter
func (r TaobaotvpaypromotioninfogetAPIRequest) GetIsTao() bool {
	return r._isTao
}
