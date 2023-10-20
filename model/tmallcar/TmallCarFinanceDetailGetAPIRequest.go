package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarfinancedetailgetAPIRequest 查询汽车金融订单信息 API请求
// tmall.car.finance.detail.get
//
// 查询汽车金融订单信息
type TmallcarfinancedetailgetAPIRequest struct {
	model.Params
	// 查询参数
	_param0 *FinanceDetailQueryReq
}

// NewTmallcarfinancedetailgetRequest 初始化TmallcarfinancedetailgetAPIRequest对象
func NewTmallcarfinancedetailgetRequest() *TmallcarfinancedetailgetAPIRequest {
	return &TmallcarfinancedetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarfinancedetailgetAPIRequest) GetApiMethodName() string {
	return "tmall.car.finance.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarfinancedetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarfinancedetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询参数
func (r *TmallcarfinancedetailgetAPIRequest) SetParam0(_param0 *FinanceDetailQueryReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallcarfinancedetailgetAPIRequest) GetParam0() *FinanceDetailQueryReq {
	return r._param0
}
