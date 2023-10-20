package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest 虚拟供应商履约接口切换 API请求
// taobao.virtual.dsf.supplier.interface.switch
//
// 虚拟供应商履约接口切换
type TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest struct {
	model.Params
	// 切流请求
	_dsfSupplierSpuRequest *DsfSupplierSpuRequest
}

// NewTaobaoVirtualDsfSupplierInterfaceSwitchRequest 初始化TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest对象
func NewTaobaoVirtualDsfSupplierInterfaceSwitchRequest() *TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest {
	return &TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest) Reset() {
	r._dsfSupplierSpuRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest) GetApiMethodName() string {
	return "taobao.virtual.dsf.supplier.interface.switch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDsfSupplierSpuRequest is DsfSupplierSpuRequest Setter
// 切流请求
func (r *TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest) SetDsfSupplierSpuRequest(_dsfSupplierSpuRequest *DsfSupplierSpuRequest) error {
	r._dsfSupplierSpuRequest = _dsfSupplierSpuRequest
	r.Set("dsf_supplier_spu_request", _dsfSupplierSpuRequest)
	return nil
}

// GetDsfSupplierSpuRequest DsfSupplierSpuRequest Getter
func (r TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest) GetDsfSupplierSpuRequest() *DsfSupplierSpuRequest {
	return r._dsfSupplierSpuRequest
}

var poolTaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVirtualDsfSupplierInterfaceSwitchRequest()
	},
}

// GetTaobaoVirtualDsfSupplierInterfaceSwitchRequest 从 sync.Pool 获取 TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest
func GetTaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest() *TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest {
	return poolTaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest.Get().(*TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest)
}

// ReleaseTaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest 将 TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest 放入 sync.Pool
func ReleaseTaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest(v *TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest) {
	v.Reset()
	poolTaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest.Put(v)
}
