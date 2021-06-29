package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销退单上传 API请求
taobao.rdc.aligenius.warehouse.reverse.uploading

主要用于商家上传仓库销退单信息
*/
type TaobaoRdcAligeniusWarehouseReverseUploadingRequest struct {
    model.Params
    // 参数
    _param0   *WarehouseReverseUploadingDto
}

// 初始化TaobaoRdcAligeniusWarehouseReverseUploadingRequest对象
func NewTaobaoRdcAligeniusWarehouseReverseUploadingRequest() *TaobaoRdcAligeniusWarehouseReverseUploadingRequest{
    return &TaobaoRdcAligeniusWarehouseReverseUploadingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseReverseUploadingRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.warehouse.reverse.uploading"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseReverseUploadingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseReverseUploadingRequest) SetParam0(_param0 *WarehouseReverseUploadingDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoRdcAligeniusWarehouseReverseUploadingRequest) GetParam0() *WarehouseReverseUploadingDto {
    return r._param0
}
