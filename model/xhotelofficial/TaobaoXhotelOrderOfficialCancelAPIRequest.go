package xhotelofficial

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfficialCancelAPIRequest 官网信用住订单取消 API请求
// taobao.xhotel.order.official.cancel
//
// 官网信用住订单取消
type TaobaoXhotelOrderOfficialCancelAPIRequest struct {
	model.Params
	// 原因描述
	_reasonText string
	// 外部订单号
	_outId string
	// 请求流水号（必须传入）
	_outUuid string
	// 暂无意义，无需传入
	_notifyUrl string
	// 淘宝订单号,必选
	_tid int64
}

// NewTaobaoXhotelOrderOfficialCancelRequest 初始化TaobaoXhotelOrderOfficialCancelAPIRequest对象
func NewTaobaoXhotelOrderOfficialCancelRequest() *TaobaoXhotelOrderOfficialCancelAPIRequest {
	return &TaobaoXhotelOrderOfficialCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.official.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReasonText is ReasonText Setter
// 原因描述
func (r *TaobaoXhotelOrderOfficialCancelAPIRequest) SetReasonText(_reasonText string) error {
	r._reasonText = _reasonText
	r.Set("reason_text", _reasonText)
	return nil
}

// GetReasonText ReasonText Getter
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetReasonText() string {
	return r._reasonText
}

// SetOutId is OutId Setter
// 外部订单号
func (r *TaobaoXhotelOrderOfficialCancelAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetOutId() string {
	return r._outId
}

// SetOutUuid is OutUuid Setter
// 请求流水号（必须传入）
func (r *TaobaoXhotelOrderOfficialCancelAPIRequest) SetOutUuid(_outUuid string) error {
	r._outUuid = _outUuid
	r.Set("out_uuid", _outUuid)
	return nil
}

// GetOutUuid OutUuid Getter
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetOutUuid() string {
	return r._outUuid
}

// SetNotifyUrl is NotifyUrl Setter
// 暂无意义，无需传入
func (r *TaobaoXhotelOrderOfficialCancelAPIRequest) SetNotifyUrl(_notifyUrl string) error {
	r._notifyUrl = _notifyUrl
	r.Set("notify_url", _notifyUrl)
	return nil
}

// GetNotifyUrl NotifyUrl Getter
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetNotifyUrl() string {
	return r._notifyUrl
}

// SetTid is Tid Setter
// 淘宝订单号,必选
func (r *TaobaoXhotelOrderOfficialCancelAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetTid() int64 {
	return r._tid
}
