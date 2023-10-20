package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrInventoryInitialAPIRequest 门店库存初始化前后端商品绑定 API请求
// tmall.nr.inventory.initial
//
// 用于门店业务的商品的初始化，前端商品和后端商品绑定，走后端库存模式
type TmallNrInventoryInitialAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *NrStoreInvItemInitialReqDto
}

// NewTmallNrInventoryInitialRequest 初始化TmallNrInventoryInitialAPIRequest对象
func NewTmallNrInventoryInitialRequest() *TmallNrInventoryInitialAPIRequest {
	return &TmallNrInventoryInitialAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrInventoryInitialAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrInventoryInitialAPIRequest) GetApiMethodName() string {
	return "tmall.nr.inventory.initial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrInventoryInitialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrInventoryInitialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求入参
func (r *TmallNrInventoryInitialAPIRequest) SetParam0(_param0 *NrStoreInvItemInitialReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallNrInventoryInitialAPIRequest) GetParam0() *NrStoreInvItemInitialReqDto {
	return r._param0
}

var poolTmallNrInventoryInitialAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrInventoryInitialRequest()
	},
}

// GetTmallNrInventoryInitialRequest 从 sync.Pool 获取 TmallNrInventoryInitialAPIRequest
func GetTmallNrInventoryInitialAPIRequest() *TmallNrInventoryInitialAPIRequest {
	return poolTmallNrInventoryInitialAPIRequest.Get().(*TmallNrInventoryInitialAPIRequest)
}

// ReleaseTmallNrInventoryInitialAPIRequest 将 TmallNrInventoryInitialAPIRequest 放入 sync.Pool
func ReleaseTmallNrInventoryInitialAPIRequest(v *TmallNrInventoryInitialAPIRequest) {
	v.Reset()
	poolTmallNrInventoryInitialAPIRequest.Put(v)
}
