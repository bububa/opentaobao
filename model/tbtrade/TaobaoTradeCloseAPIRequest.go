package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradecloseAPIRequest 卖家关闭一笔交易 API请求
// taobao.trade.close
//
// 关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。
type TaobaotradecloseAPIRequest struct {
	model.Params
	// 交易关闭原因。可以选择的理由有：1.未及时付款2、买家不想买了3、买家信息填写错误，重新拍4、恶意买家/同行捣乱5、缺货6、买家拍错了7、同城见面交易
	_closeReason string
	// 主订单或子订单编号。
	_tid int64
}

// NewTaobaotradecloseRequest 初始化TaobaotradecloseAPIRequest对象
func NewTaobaotradecloseRequest() *TaobaotradecloseAPIRequest {
	return &TaobaotradecloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradecloseAPIRequest) GetApiMethodName() string {
	return "taobao.trade.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradecloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradecloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCloseReason is CloseReason Setter
// 交易关闭原因。可以选择的理由有：1.未及时付款2、买家不想买了3、买家信息填写错误，重新拍4、恶意买家/同行捣乱5、缺货6、买家拍错了7、同城见面交易
func (r *TaobaotradecloseAPIRequest) SetCloseReason(_closeReason string) error {
	r._closeReason = _closeReason
	r.Set("close_reason", _closeReason)
	return nil
}

// GetCloseReason CloseReason Getter
func (r TaobaotradecloseAPIRequest) GetCloseReason() string {
	return r._closeReason
}

// SetTid is Tid Setter
// 主订单或子订单编号。
func (r *TaobaotradecloseAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradecloseAPIRequest) GetTid() int64 {
	return r._tid
}
