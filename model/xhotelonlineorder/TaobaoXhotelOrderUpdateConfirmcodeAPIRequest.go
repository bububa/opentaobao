package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderUpdateConfirmcodeAPIRequest 推送及更新订单确认号 API请求
// taobao.xhotel.order.update.confirmcode
//
// 商家拿到订单确认号后，异步推送给飞猪或更新给飞猪。订单确认号用于到店查无时的紧急查单依据。
type TaobaoXhotelOrderUpdateConfirmcodeAPIRequest struct {
	model.Params
	// 系统入参
	_param *UpdateOrderConfirmCodeParam
}

// NewTaobaoXhotelOrderUpdateConfirmcodeRequest 初始化TaobaoXhotelOrderUpdateConfirmcodeAPIRequest对象
func NewTaobaoXhotelOrderUpdateConfirmcodeRequest() *TaobaoXhotelOrderUpdateConfirmcodeAPIRequest {
	return &TaobaoXhotelOrderUpdateConfirmcodeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelOrderUpdateConfirmcodeAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderUpdateConfirmcodeAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.update.confirmcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderUpdateConfirmcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelOrderUpdateConfirmcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 系统入参
func (r *TaobaoXhotelOrderUpdateConfirmcodeAPIRequest) SetParam(_param *UpdateOrderConfirmCodeParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoXhotelOrderUpdateConfirmcodeAPIRequest) GetParam() *UpdateOrderConfirmCodeParam {
	return r._param
}

var poolTaobaoXhotelOrderUpdateConfirmcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelOrderUpdateConfirmcodeRequest()
	},
}

// GetTaobaoXhotelOrderUpdateConfirmcodeRequest 从 sync.Pool 获取 TaobaoXhotelOrderUpdateConfirmcodeAPIRequest
func GetTaobaoXhotelOrderUpdateConfirmcodeAPIRequest() *TaobaoXhotelOrderUpdateConfirmcodeAPIRequest {
	return poolTaobaoXhotelOrderUpdateConfirmcodeAPIRequest.Get().(*TaobaoXhotelOrderUpdateConfirmcodeAPIRequest)
}

// ReleaseTaobaoXhotelOrderUpdateConfirmcodeAPIRequest 将 TaobaoXhotelOrderUpdateConfirmcodeAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelOrderUpdateConfirmcodeAPIRequest(v *TaobaoXhotelOrderUpdateConfirmcodeAPIRequest) {
	v.Reset()
	poolTaobaoXhotelOrderUpdateConfirmcodeAPIRequest.Put(v)
}
