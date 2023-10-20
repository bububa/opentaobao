package axintrade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripAxinTransFundConfirmAPIRequest) Reset() {
	r._outerOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransFundConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.fund.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransFundConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripAxinTransFundConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoAlitripAxinTransFundConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripAxinTransFundConfirmRequest()
	},
}

// GetTaobaoAlitripAxinTransFundConfirmRequest 从 sync.Pool 获取 TaobaoAlitripAxinTransFundConfirmAPIRequest
func GetTaobaoAlitripAxinTransFundConfirmAPIRequest() *TaobaoAlitripAxinTransFundConfirmAPIRequest {
	return poolTaobaoAlitripAxinTransFundConfirmAPIRequest.Get().(*TaobaoAlitripAxinTransFundConfirmAPIRequest)
}

// ReleaseTaobaoAlitripAxinTransFundConfirmAPIRequest 将 TaobaoAlitripAxinTransFundConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripAxinTransFundConfirmAPIRequest(v *TaobaoAlitripAxinTransFundConfirmAPIRequest) {
	v.Reset()
	poolTaobaoAlitripAxinTransFundConfirmAPIRequest.Put(v)
}
