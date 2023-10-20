package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSdtstatusAPIRequest 菜鸟裹裹运单状态查询 API请求
// taobao.omniorder.store.sdtstatus
//
// 提供给商家查询运力单的状态。
type TaobaoOmniorderStoreSdtstatusAPIRequest struct {
	model.Params
	// 菜鸟裹裹的包裹ID
	_packageId int64
}

// NewTaobaoOmniorderStoreSdtstatusRequest 初始化TaobaoOmniorderStoreSdtstatusAPIRequest对象
func NewTaobaoOmniorderStoreSdtstatusRequest() *TaobaoOmniorderStoreSdtstatusAPIRequest {
	return &TaobaoOmniorderStoreSdtstatusAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderStoreSdtstatusAPIRequest) Reset() {
	r._packageId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSdtstatusAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.sdtstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSdtstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStoreSdtstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPackageId is PackageId Setter
// 菜鸟裹裹的包裹ID
func (r *TaobaoOmniorderStoreSdtstatusAPIRequest) SetPackageId(_packageId int64) error {
	r._packageId = _packageId
	r.Set("package_id", _packageId)
	return nil
}

// GetPackageId PackageId Getter
func (r TaobaoOmniorderStoreSdtstatusAPIRequest) GetPackageId() int64 {
	return r._packageId
}

var poolTaobaoOmniorderStoreSdtstatusAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderStoreSdtstatusRequest()
	},
}

// GetTaobaoOmniorderStoreSdtstatusRequest 从 sync.Pool 获取 TaobaoOmniorderStoreSdtstatusAPIRequest
func GetTaobaoOmniorderStoreSdtstatusAPIRequest() *TaobaoOmniorderStoreSdtstatusAPIRequest {
	return poolTaobaoOmniorderStoreSdtstatusAPIRequest.Get().(*TaobaoOmniorderStoreSdtstatusAPIRequest)
}

// ReleaseTaobaoOmniorderStoreSdtstatusAPIRequest 将 TaobaoOmniorderStoreSdtstatusAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderStoreSdtstatusAPIRequest(v *TaobaoOmniorderStoreSdtstatusAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderStoreSdtstatusAPIRequest.Put(v)
}
