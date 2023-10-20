package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstoresdtstatusAPIRequest 菜鸟裹裹运单状态查询 API请求
// taobao.omniorder.store.sdtstatus
//
// 提供给商家查询运力单的状态。
type TaobaoomniorderstoresdtstatusAPIRequest struct {
	model.Params
	// 菜鸟裹裹的包裹ID
	_packageId int64
}

// NewTaobaoomniorderstoresdtstatusRequest 初始化TaobaoomniorderstoresdtstatusAPIRequest对象
func NewTaobaoomniorderstoresdtstatusRequest() *TaobaoomniorderstoresdtstatusAPIRequest {
	return &TaobaoomniorderstoresdtstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstoresdtstatusAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.sdtstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstoresdtstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstoresdtstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPackageId is PackageId Setter
// 菜鸟裹裹的包裹ID
func (r *TaobaoomniorderstoresdtstatusAPIRequest) SetPackageId(_packageId int64) error {
	r._packageId = _packageId
	r.Set("package_id", _packageId)
	return nil
}

// GetPackageId PackageId Getter
func (r TaobaoomniorderstoresdtstatusAPIRequest) GetPackageId() int64 {
	return r._packageId
}
