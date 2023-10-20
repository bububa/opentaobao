package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrInventoryUpdateAPIRequest 门店业务同步库存 API请求
// tmall.nr.inventory.update
//
// 用于商家每日同步更新门店库存
type TmallNrInventoryUpdateAPIRequest struct {
	model.Params
	// 入参
	_param0 *NrInventoryUpdateReqDto
}

// NewTmallNrInventoryUpdateRequest 初始化TmallNrInventoryUpdateAPIRequest对象
func NewTmallNrInventoryUpdateRequest() *TmallNrInventoryUpdateAPIRequest {
	return &TmallNrInventoryUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrInventoryUpdateAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrInventoryUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.nr.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrInventoryUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrInventoryUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TmallNrInventoryUpdateAPIRequest) SetParam0(_param0 *NrInventoryUpdateReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallNrInventoryUpdateAPIRequest) GetParam0() *NrInventoryUpdateReqDto {
	return r._param0
}

var poolTmallNrInventoryUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrInventoryUpdateRequest()
	},
}

// GetTmallNrInventoryUpdateRequest 从 sync.Pool 获取 TmallNrInventoryUpdateAPIRequest
func GetTmallNrInventoryUpdateAPIRequest() *TmallNrInventoryUpdateAPIRequest {
	return poolTmallNrInventoryUpdateAPIRequest.Get().(*TmallNrInventoryUpdateAPIRequest)
}

// ReleaseTmallNrInventoryUpdateAPIRequest 将 TmallNrInventoryUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallNrInventoryUpdateAPIRequest(v *TmallNrInventoryUpdateAPIRequest) {
	v.Reset()
	poolTmallNrInventoryUpdateAPIRequest.Put(v)
}
