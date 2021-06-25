package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销退单上传 APIRequest
taobao.rdc.aligenius.warehouse.reverse.uploading

主要用于商家上传仓库销退单信息
*/
type TaobaoRdcAligeniusWarehouseReverseUploadingRequest struct {
    model.Params

    // 参数
    param0   *WarehouseReverseUploadingDto 

}

func NewTaobaoRdcAligeniusWarehouseReverseUploadingRequest() *TaobaoRdcAligeniusWarehouseReverseUploadingRequest{
    return &TaobaoRdcAligeniusWarehouseReverseUploadingRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdcAligeniusWarehouseReverseUploadingRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.warehouse.reverse.uploading"
}

func (r TaobaoRdcAligeniusWarehouseReverseUploadingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdcAligeniusWarehouseReverseUploadingRequest) SetParam0(param0 *WarehouseReverseUploadingDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoRdcAligeniusWarehouseReverseUploadingRequest) GetParam0() *WarehouseReverseUploadingDto {
    return r.param0
}

