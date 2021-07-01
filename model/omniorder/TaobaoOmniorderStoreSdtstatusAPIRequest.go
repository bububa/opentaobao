package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreSdtstatusAPIRequest
菜鸟裹裹运单状态查询 API请求
taobao.omniorder.store.sdtstatus

提供给商家查询运力单的状态。 */
type TaobaoOmniorderStoreSdtstatusAPIRequest struct {
	model.Params
	// 菜鸟裹裹的包裹ID
	_packageId int64
}

// NewTaobaoOmniorderStoreSdtstatusRequest 初始化TaobaoOmniorderStoreSdtstatusAPIRequest对象
func NewTaobaoOmniorderStoreSdtstatusRequest() *TaobaoOmniorderStoreSdtstatusAPIRequest {
	return &TaobaoOmniorderStoreSdtstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSdtstatusAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.sdtstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSdtstatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PackageId Setter
// 菜鸟裹裹的包裹ID
func (r *TaobaoOmniorderStoreSdtstatusAPIRequest) SetPackageId(_packageId int64) error {
	r._packageId = _packageId
	r.Set("package_id", _packageId)
	return nil
}

// Get PackageId Getter
func (r TaobaoOmniorderStoreSdtstatusAPIRequest) GetPackageId() int64 {
	return r._packageId
}
