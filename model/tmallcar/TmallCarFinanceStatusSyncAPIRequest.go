package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarfinancestatussyncAPIRequest 汽车金融状态同步 API请求
// tmall.car.finance.status.sync
//
// 汽车金融状态同步
type TmallcarfinancestatussyncAPIRequest struct {
	model.Params
	// 状态同步数据
	_param0 *CreditLoanStatusSyncReq
}

// NewTmallcarfinancestatussyncRequest 初始化TmallcarfinancestatussyncAPIRequest对象
func NewTmallcarfinancestatussyncRequest() *TmallcarfinancestatussyncAPIRequest {
	return &TmallcarfinancestatussyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarfinancestatussyncAPIRequest) GetApiMethodName() string {
	return "tmall.car.finance.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarfinancestatussyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarfinancestatussyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 状态同步数据
func (r *TmallcarfinancestatussyncAPIRequest) SetParam0(_param0 *CreditLoanStatusSyncReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallcarfinancestatussyncAPIRequest) GetParam0() *CreditLoanStatusSyncReq {
	return r._param0
}
