package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrinventoryupdateAPIRequest 门店业务同步库存 API请求
// tmall.nr.inventory.update
//
// 用于商家每日同步更新门店库存
type TmallnrinventoryupdateAPIRequest struct {
	model.Params
	// 入参
	_param0 *NrInventoryUpdateReqDto
}

// NewTmallnrinventoryupdateRequest 初始化TmallnrinventoryupdateAPIRequest对象
func NewTmallnrinventoryupdateRequest() *TmallnrinventoryupdateAPIRequest {
	return &TmallnrinventoryupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrinventoryupdateAPIRequest) GetApiMethodName() string {
	return "tmall.nr.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrinventoryupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrinventoryupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TmallnrinventoryupdateAPIRequest) SetParam0(_param0 *NrInventoryUpdateReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallnrinventoryupdateAPIRequest) GetParam0() *NrInventoryUpdateReqDto {
	return r._param0
}
