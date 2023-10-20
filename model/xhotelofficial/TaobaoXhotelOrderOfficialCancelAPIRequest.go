package xhotelofficial

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelOrderOfficialCancelAPIRequest) Reset() {
	r._reasonText = ""
	r._outId = ""
	r._outUuid = ""
	r._notifyUrl = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.official.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelOrderOfficialCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoXhotelOrderOfficialCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelOrderOfficialCancelRequest()
	},
}

// GetTaobaoXhotelOrderOfficialCancelRequest 从 sync.Pool 获取 TaobaoXhotelOrderOfficialCancelAPIRequest
func GetTaobaoXhotelOrderOfficialCancelAPIRequest() *TaobaoXhotelOrderOfficialCancelAPIRequest {
	return poolTaobaoXhotelOrderOfficialCancelAPIRequest.Get().(*TaobaoXhotelOrderOfficialCancelAPIRequest)
}

// ReleaseTaobaoXhotelOrderOfficialCancelAPIRequest 将 TaobaoXhotelOrderOfficialCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelOrderOfficialCancelAPIRequest(v *TaobaoXhotelOrderOfficialCancelAPIRequest) {
	v.Reset()
	poolTaobaoXhotelOrderOfficialCancelAPIRequest.Put(v)
}
