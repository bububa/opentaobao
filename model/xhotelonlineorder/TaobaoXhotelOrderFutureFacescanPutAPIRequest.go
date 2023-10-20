package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderFutureFacescanPutAPIRequest 未来酒店扫脸信息上传 API请求
// taobao.xhotel.order.future.facescan.put
//
// 未来酒店扫脸信息上传服务，用于悉尔等厂商的扫脸设备对接
type TaobaoXhotelOrderFutureFacescanPutAPIRequest struct {
	model.Params
	// 扫脸参数
	_faceScanParam *FaceScanParam
}

// NewTaobaoXhotelOrderFutureFacescanPutRequest 初始化TaobaoXhotelOrderFutureFacescanPutAPIRequest对象
func NewTaobaoXhotelOrderFutureFacescanPutRequest() *TaobaoXhotelOrderFutureFacescanPutAPIRequest {
	return &TaobaoXhotelOrderFutureFacescanPutAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelOrderFutureFacescanPutAPIRequest) Reset() {
	r._faceScanParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderFutureFacescanPutAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.future.facescan.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderFutureFacescanPutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelOrderFutureFacescanPutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFaceScanParam is FaceScanParam Setter
// 扫脸参数
func (r *TaobaoXhotelOrderFutureFacescanPutAPIRequest) SetFaceScanParam(_faceScanParam *FaceScanParam) error {
	r._faceScanParam = _faceScanParam
	r.Set("face_scan_param", _faceScanParam)
	return nil
}

// GetFaceScanParam FaceScanParam Getter
func (r TaobaoXhotelOrderFutureFacescanPutAPIRequest) GetFaceScanParam() *FaceScanParam {
	return r._faceScanParam
}

var poolTaobaoXhotelOrderFutureFacescanPutAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelOrderFutureFacescanPutRequest()
	},
}

// GetTaobaoXhotelOrderFutureFacescanPutRequest 从 sync.Pool 获取 TaobaoXhotelOrderFutureFacescanPutAPIRequest
func GetTaobaoXhotelOrderFutureFacescanPutAPIRequest() *TaobaoXhotelOrderFutureFacescanPutAPIRequest {
	return poolTaobaoXhotelOrderFutureFacescanPutAPIRequest.Get().(*TaobaoXhotelOrderFutureFacescanPutAPIRequest)
}

// ReleaseTaobaoXhotelOrderFutureFacescanPutAPIRequest 将 TaobaoXhotelOrderFutureFacescanPutAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelOrderFutureFacescanPutAPIRequest(v *TaobaoXhotelOrderFutureFacescanPutAPIRequest) {
	v.Reset()
	poolTaobaoXhotelOrderFutureFacescanPutAPIRequest.Put(v)
}
