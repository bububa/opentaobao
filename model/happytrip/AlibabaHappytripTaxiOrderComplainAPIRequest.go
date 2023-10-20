package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiordercomplainAPIRequest 用户投诉 API请求
// alibaba.happytrip.taxi.order.complain
//
// 一个订单只能投诉一次，不可重复投诉
//
// 投诉选项
// 0	其他原因；
// 1	司机原因导致行程被取消；
// 2	服务态度恶劣；
// 3	未坐车产生费用；
// 4	额外收取不合理费用；
// 5	路不熟多产生费用；
// 6	提前计费；
// 7	未及时结束计费；
// 8	司机绕路；
// 9	司机迟到；
// 10	司机爽约或拒载；
// 11	骚扰乘客；
// 12	危险驾驶；
// 13	不是订单显示车辆或司机；
type AlibabahappytriptaxiordercomplainAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 投诉选项外的其他投诉内容,不能多于40个汉字
	_content string
	// 该用户的真实号码
	_mobile string
	// 投诉选项 0	其他原因； 1	司机原因导致行程被取消； 2	服务态度恶劣；	 3	未坐车产生费用；	 4	额外收取不合理费用； 5	路不熟多产生费用；	 6	提前计费； 7	未及时结束计费；	 8	司机绕路；	 9	司机迟到；	 10	司机爽约或拒载；	 11	骚扰乘客；	 12	危险驾驶；	 13	不是订单显示车辆或司机；
	_type int64
	// 投诉时间（毫秒）
	_time int64
}

// NewAlibabahappytriptaxiordercomplainRequest 初始化AlibabahappytriptaxiordercomplainAPIRequest对象
func NewAlibabahappytriptaxiordercomplainRequest() *AlibabahappytriptaxiordercomplainAPIRequest {
	return &AlibabahappytriptaxiordercomplainAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxiordercomplainAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.complain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxiordercomplainAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxiordercomplainAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabahappytriptaxiordercomplainAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxiordercomplainAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetContent is Content Setter
// 投诉选项外的其他投诉内容,不能多于40个汉字
func (r *AlibabahappytriptaxiordercomplainAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabahappytriptaxiordercomplainAPIRequest) GetContent() string {
	return r._content
}

// SetMobile is Mobile Setter
// 该用户的真实号码
func (r *AlibabahappytriptaxiordercomplainAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabahappytriptaxiordercomplainAPIRequest) GetMobile() string {
	return r._mobile
}

// SetType is Type Setter
// 投诉选项 0	其他原因； 1	司机原因导致行程被取消； 2	服务态度恶劣；	 3	未坐车产生费用；	 4	额外收取不合理费用； 5	路不熟多产生费用；	 6	提前计费； 7	未及时结束计费；	 8	司机绕路；	 9	司机迟到；	 10	司机爽约或拒载；	 11	骚扰乘客；	 12	危险驾驶；	 13	不是订单显示车辆或司机；
func (r *AlibabahappytriptaxiordercomplainAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabahappytriptaxiordercomplainAPIRequest) GetType() int64 {
	return r._type
}

// SetTime is Time Setter
// 投诉时间（毫秒）
func (r *AlibabahappytriptaxiordercomplainAPIRequest) SetTime(_time int64) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r AlibabahappytriptaxiordercomplainAPIRequest) GetTime() int64 {
	return r._time
}
