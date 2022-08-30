package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayPromotionInfoGetAPIRequest tv支付查询消费抽奖配置 API请求
// taobao.tvpay.promotion.info.get
//
// 查询消费抽奖配置
type TaobaoTvpayPromotionInfoGetAPIRequest struct {
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

// NewTaobaoTvpayPromotionInfoGetRequest 初始化TaobaoTvpayPromotionInfoGetAPIRequest对象
func NewTaobaoTvpayPromotionInfoGetRequest() *TaobaoTvpayPromotionInfoGetAPIRequest {
	return &TaobaoTvpayPromotionInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayPromotionInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.promotion.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayPromotionInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoTvpayPromotionInfoGetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoTvpayPromotionInfoGetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaoTvpayPromotionInfoGetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaoTvpayPromotionInfoGetAPIRequest) GetFrom() string {
	return r._from
}

// SetExtOrderId is ExtOrderId Setter
// 淘系订单号
func (r *TaobaoTvpayPromotionInfoGetAPIRequest) SetExtOrderId(_extOrderId string) error {
	r._extOrderId = _extOrderId
	r.Set("ext_order_id", _extOrderId)
	return nil
}

// GetExtOrderId ExtOrderId Getter
func (r TaobaoTvpayPromotionInfoGetAPIRequest) GetExtOrderId() string {
	return r._extOrderId
}

// SetSubject is Subject Setter
// 标题
func (r *TaobaoTvpayPromotionInfoGetAPIRequest) SetSubject(_subject string) error {
	r._subject = _subject
	r.Set("subject", _subject)
	return nil
}

// GetSubject Subject Getter
func (r TaobaoTvpayPromotionInfoGetAPIRequest) GetSubject() string {
	return r._subject
}

// SetSubjectId is SubjectId Setter
// 商品id
func (r *TaobaoTvpayPromotionInfoGetAPIRequest) SetSubjectId(_subjectId string) error {
	r._subjectId = _subjectId
	r.Set("subject_id", _subjectId)
	return nil
}

// GetSubjectId SubjectId Getter
func (r TaobaoTvpayPromotionInfoGetAPIRequest) GetSubjectId() string {
	return r._subjectId
}

// SetIsTao is IsTao Setter
// 是否淘系
func (r *TaobaoTvpayPromotionInfoGetAPIRequest) SetIsTao(_isTao bool) error {
	r._isTao = _isTao
	r.Set("is_tao", _isTao)
	return nil
}

// GetIsTao IsTao Getter
func (r TaobaoTvpayPromotionInfoGetAPIRequest) GetIsTao() bool {
	return r._isTao
}
