package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口仓库悬挂链信息上报 API请求
taobao.wdk.equipment.conveyor.info.upload

五道口仓库悬挂链信息上传
*/
type TaobaoWdkEquipmentConveyorInfoUploadRequest struct {
    model.Params
    // 上传信息
    param0   string
}

// 初始化TaobaoWdkEquipmentConveyorInfoUploadRequest对象
func NewTaobaoWdkEquipmentConveyorInfoUploadRequest() *TaobaoWdkEquipmentConveyorInfoUploadRequest{
    return &TaobaoWdkEquipmentConveyorInfoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorInfoUploadRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.info.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorInfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 上传信息
func (r *TaobaoWdkEquipmentConveyorInfoUploadRequest) SetParam0(param0 string) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoWdkEquipmentConveyorInfoUploadRequest) GetParam0() string {
    return r.param0
}
