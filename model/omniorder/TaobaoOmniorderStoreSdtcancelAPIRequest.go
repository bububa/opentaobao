package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSdtcancelAPIRequest 通知速店通取消取号 API请求
// taobao.omniorder.store.sdtcancel
//
// 通知速店通取消取号
type TaobaoOmniorderStoreSdtcancelAPIRequest struct {
	model.Params
	// 取号返回的packageId
	_packageId int64
}

// NewTaobaoOmniorderStoreSdtcancelRequest 初始化TaobaoOmniorderStoreSdtcancelAPIRequest对象
func NewTaobaoOmniorderStoreSdtcancelRequest() *TaobaoOmniorderStoreSdtcancelAPIRequest {
	return &TaobaoOmniorderStoreSdtcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSdtcancelAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.sdtcancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSdtcancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PackageId Setter
// 取号返回的packageId
func (r *TaobaoOmniorderStoreSdtcancelAPIRequest) SetPackageId(_packageId int64) error {
	r._packageId = _packageId
	r.Set("package_id", _packageId)
	return nil
}

// Get PackageId Getter
func (r TaobaoOmniorderStoreSdtcancelAPIRequest) GetPackageId() int64 {
	return r._packageId
}
