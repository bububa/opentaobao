package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMallitemcenterSupplierAbilityUpdateAPIRequest 门店服务能力授权接口 API请求
// tmall.mallitemcenter.supplier.ability.update
//
// 门店服务能力授权
type TmallMallitemcenterSupplierAbilityUpdateAPIRequest struct {
	model.Params
	// 入参
	_param0 *EnableServiceStoreRequestDto
}

// NewTmallMallitemcenterSupplierAbilityUpdateRequest 初始化TmallMallitemcenterSupplierAbilityUpdateAPIRequest对象
func NewTmallMallitemcenterSupplierAbilityUpdateRequest() *TmallMallitemcenterSupplierAbilityUpdateAPIRequest {
	return &TmallMallitemcenterSupplierAbilityUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMallitemcenterSupplierAbilityUpdateAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterSupplierAbilityUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.mallitemcenter.supplier.ability.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterSupplierAbilityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMallitemcenterSupplierAbilityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TmallMallitemcenterSupplierAbilityUpdateAPIRequest) SetParam0(_param0 *EnableServiceStoreRequestDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallMallitemcenterSupplierAbilityUpdateAPIRequest) GetParam0() *EnableServiceStoreRequestDto {
	return r._param0
}

var poolTmallMallitemcenterSupplierAbilityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMallitemcenterSupplierAbilityUpdateRequest()
	},
}

// GetTmallMallitemcenterSupplierAbilityUpdateRequest 从 sync.Pool 获取 TmallMallitemcenterSupplierAbilityUpdateAPIRequest
func GetTmallMallitemcenterSupplierAbilityUpdateAPIRequest() *TmallMallitemcenterSupplierAbilityUpdateAPIRequest {
	return poolTmallMallitemcenterSupplierAbilityUpdateAPIRequest.Get().(*TmallMallitemcenterSupplierAbilityUpdateAPIRequest)
}

// ReleaseTmallMallitemcenterSupplierAbilityUpdateAPIRequest 将 TmallMallitemcenterSupplierAbilityUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallMallitemcenterSupplierAbilityUpdateAPIRequest(v *TmallMallitemcenterSupplierAbilityUpdateAPIRequest) {
	v.Reset()
	poolTmallMallitemcenterSupplierAbilityUpdateAPIRequest.Put(v)
}
