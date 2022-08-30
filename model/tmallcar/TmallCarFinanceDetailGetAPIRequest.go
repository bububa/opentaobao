package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarFinanceDetailGetAPIRequest 查询汽车金融订单信息 API请求
// tmall.car.finance.detail.get
//
// 查询汽车金融订单信息
type TmallCarFinanceDetailGetAPIRequest struct {
	model.Params
	// 查询参数
	_param0 *FinanceDetailQueryReq
}

// NewTmallCarFinanceDetailGetRequest 初始化TmallCarFinanceDetailGetAPIRequest对象
func NewTmallCarFinanceDetailGetRequest() *TmallCarFinanceDetailGetAPIRequest {
	return &TmallCarFinanceDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarFinanceDetailGetAPIRequest) GetApiMethodName() string {
	return "tmall.car.finance.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarFinanceDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 查询参数
func (r *TmallCarFinanceDetailGetAPIRequest) SetParam0(_param0 *FinanceDetailQueryReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallCarFinanceDetailGetAPIRequest) GetParam0() *FinanceDetailQueryReq {
	return r._param0
}
