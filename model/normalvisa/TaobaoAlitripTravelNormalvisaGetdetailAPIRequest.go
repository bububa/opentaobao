package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelnormalvisagetdetailAPIRequest 获取单笔订单的详情 API请求
// taobao.alitrip.travel.normalvisa.getdetail
//
// 获取单笔签证的详细记录
type TaobaoalitriptravelnormalvisagetdetailAPIRequest struct {
	model.Params
	// 订单id
	_bizOrderId int64
}

// NewTaobaoalitriptravelnormalvisagetdetailRequest 初始化TaobaoalitriptravelnormalvisagetdetailAPIRequest对象
func NewTaobaoalitriptravelnormalvisagetdetailRequest() *TaobaoalitriptravelnormalvisagetdetailAPIRequest {
	return &TaobaoalitriptravelnormalvisagetdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelnormalvisagetdetailAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.getdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelnormalvisagetdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelnormalvisagetdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单id
func (r *TaobaoalitriptravelnormalvisagetdetailAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoalitriptravelnormalvisagetdetailAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
