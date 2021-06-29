package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
未来酒店扫脸信息上传 API请求
taobao.xhotel.order.future.facescan.put

未来酒店扫脸信息上传服务，用于悉尔等厂商的扫脸设备对接
*/
type TaobaoXhotelOrderFutureFacescanPutRequest struct {
    model.Params
    // 扫脸参数
    _faceScanParam   *FaceScanParam
}

// 初始化TaobaoXhotelOrderFutureFacescanPutRequest对象
func NewTaobaoXhotelOrderFutureFacescanPutRequest() *TaobaoXhotelOrderFutureFacescanPutRequest{
    return &TaobaoXhotelOrderFutureFacescanPutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderFutureFacescanPutRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.future.facescan.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderFutureFacescanPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FaceScanParam Setter
// 扫脸参数
func (r *TaobaoXhotelOrderFutureFacescanPutRequest) SetFaceScanParam(_faceScanParam *FaceScanParam) error {
    r._faceScanParam = _faceScanParam
    r.Set("face_scan_param", _faceScanParam)
    return nil
}

// FaceScanParam Getter
func (r TaobaoXhotelOrderFutureFacescanPutRequest) GetFaceScanParam() *FaceScanParam {
    return r._faceScanParam
}
