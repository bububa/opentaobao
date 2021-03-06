package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransFundConfirmAPIRequest 确认资金单 API请求
// taobao.alitrip.axin.trans.fund.confirm
//
// 通过外部订单号进行资金结算
type TaobaoAlitripAxinTransFundConfirmAPIRequest struct {
	model.Params
	// 外部订单编号
	_outerOrderId string
}

// NewTaobaoAlitripAxinTransFundConfirmRequest 初始化TaobaoAlitripAxinTransFundConfirmAPIRequest对象
func NewTaobaoAlitripAxinTransFundConfirmRequest() *TaobaoAlitripAxinTransFundConfirmAPIRequest {
	return &TaobaoAlitripAxinTransFundConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransFundConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.fund.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransFundConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterOrderId is OuterOrderId Setter
// 外部订单编号
func (r *TaobaoAlitripAxinTransFundConfirmAPIRequest) SetOuterOrderId(_outerOrderId string) error {
	r._outerOrderId = _outerOrderId
	r.Set("outer_order_id", _outerOrderId)
	return nil
}

// GetOuterOrderId OuterOrderId Getter
func (r TaobaoAlitripAxinTransFundConfirmAPIRequest) GetOuterOrderId() string {
	return r._outerOrderId
}
