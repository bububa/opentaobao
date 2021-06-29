package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
未来酒店扫脸信息上传 APIRequest
taobao.xhotel.order.future.facescan.put

未来酒店扫脸信息上传服务，用于悉尔等厂商的扫脸设备对接
*/
type TaobaoXhotelOrderFutureFacescanPutRequest struct {
    model.Params

    // 扫脸参数
    faceScanParam   *FaceScanParam 

}

func NewTaobaoXhotelOrderFutureFacescanPutRequest() *TaobaoXhotelOrderFutureFacescanPutRequest{
    return &TaobaoXhotelOrderFutureFacescanPutRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderFutureFacescanPutRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.future.facescan.put"
}

func (r TaobaoXhotelOrderFutureFacescanPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderFutureFacescanPutRequest) SetFaceScanParam(faceScanParam *FaceScanParam) error {
    r.faceScanParam = faceScanParam
    r.Set("face_scan_param", faceScanParam)
    return nil
}

func (r TaobaoXhotelOrderFutureFacescanPutRequest) GetFaceScanParam() *FaceScanParam {
    return r.faceScanParam
}

