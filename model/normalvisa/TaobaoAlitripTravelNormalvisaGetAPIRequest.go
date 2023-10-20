package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelnormalvisagetAPIRequest 获取签证记录 API请求
// taobao.alitrip.travel.normalvisa.get
//
// 用于获取普通签证的记录信息
type TaobaoalitriptravelnormalvisagetAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewTaobaoalitriptravelnormalvisagetRequest 初始化TaobaoalitriptravelnormalvisagetAPIRequest对象
func NewTaobaoalitriptravelnormalvisagetRequest() *TaobaoalitriptravelnormalvisagetAPIRequest {
	return &TaobaoalitriptravelnormalvisagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelnormalvisagetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelnormalvisagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelnormalvisagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaoalitriptravelnormalvisagetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoalitriptravelnormalvisagetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
