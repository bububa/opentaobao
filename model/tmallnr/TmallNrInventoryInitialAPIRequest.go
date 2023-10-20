package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrinventoryinitialAPIRequest 门店库存初始化前后端商品绑定 API请求
// tmall.nr.inventory.initial
//
// 用于门店业务的商品的初始化，前端商品和后端商品绑定，走后端库存模式
type TmallnrinventoryinitialAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *NrStoreInvItemInitialReqDto
}

// NewTmallnrinventoryinitialRequest 初始化TmallnrinventoryinitialAPIRequest对象
func NewTmallnrinventoryinitialRequest() *TmallnrinventoryinitialAPIRequest {
	return &TmallnrinventoryinitialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrinventoryinitialAPIRequest) GetApiMethodName() string {
	return "tmall.nr.inventory.initial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrinventoryinitialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrinventoryinitialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求入参
func (r *TmallnrinventoryinitialAPIRequest) SetParam0(_param0 *NrStoreInvItemInitialReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallnrinventoryinitialAPIRequest) GetParam0() *NrStoreInvItemInitialReqDto {
	return r._param0
}
