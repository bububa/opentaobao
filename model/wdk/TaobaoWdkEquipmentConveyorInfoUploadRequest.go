package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口仓库悬挂链信息上报 APIRequest
taobao.wdk.equipment.conveyor.info.upload

五道口仓库悬挂链信息上传
*/
type TaobaoWdkEquipmentConveyorInfoUploadRequest struct {
    model.Params

    // 上传信息
    param0   string 

}

func NewTaobaoWdkEquipmentConveyorInfoUploadRequest() *TaobaoWdkEquipmentConveyorInfoUploadRequest{
    return &TaobaoWdkEquipmentConveyorInfoUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentConveyorInfoUploadRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.info.upload"
}

func (r TaobaoWdkEquipmentConveyorInfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentConveyorInfoUploadRequest) SetParam0(param0 string) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoWdkEquipmentConveyorInfoUploadRequest) GetParam0() string {
    return r.param0
}

