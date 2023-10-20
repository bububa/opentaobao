package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillSoldOrderlistQueryAPIRequest 零售商获取品牌商的特定订单列表 API请求
// tmall.nr.fulfill.sold.orderlist.query
//
// 零售商获取品牌商的特定订单列表，后端服务有零售商和品牌商的绑定关系，存在开关控制；同时后端存在定时送业务等特殊业务的校验，非同城配送业务不能返回，返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位
type TmallNrFulfillSoldOrderlistQueryAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *NrTimingOrderSoldQueryReqDto
}

// NewTmallNrFulfillSoldOrderlistQueryRequest 初始化TmallNrFulfillSoldOrderlistQueryAPIRequest对象
func NewTmallNrFulfillSoldOrderlistQueryRequest() *TmallNrFulfillSoldOrderlistQueryAPIRequest {
	return &TmallNrFulfillSoldOrderlistQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrFulfillSoldOrderlistQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nr.fulfill.sold.orderlist.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrFulfillSoldOrderlistQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrFulfillSoldOrderlistQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参对象
func (r *TmallNrFulfillSoldOrderlistQueryAPIRequest) SetParam0(_param0 *NrTimingOrderSoldQueryReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallNrFulfillSoldOrderlistQueryAPIRequest) GetParam0() *NrTimingOrderSoldQueryReqDto {
	return r._param0
}
