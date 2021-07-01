package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrInventoryUpdateAPIRequest
门店业务同步库存 API请求
tmall.nr.inventory.update

用于商家每日同步更新门店库存 */
type TmallNrInventoryUpdateAPIRequest struct {
	model.Params
	// 入参
	_param0 *NrInventoryUpdateReqDto
}

// NewTmallNrInventoryUpdateRequest 初始化TmallNrInventoryUpdateAPIRequest对象
func NewTmallNrInventoryUpdateRequest() *TmallNrInventoryUpdateAPIRequest {
	return &TmallNrInventoryUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrInventoryUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.nr.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrInventoryUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 入参
func (r *TmallNrInventoryUpdateAPIRequest) SetParam0(_param0 *NrInventoryUpdateReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TmallNrInventoryUpdateAPIRequest) GetParam0() *NrInventoryUpdateReqDto {
	return r._param0
}
