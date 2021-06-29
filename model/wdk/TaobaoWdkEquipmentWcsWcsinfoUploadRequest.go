package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链业务信息上传 API请求
taobao.wdk.equipment.wcs.wcsinfo.upload

五道口仓库悬挂链信息上传
*/
type TaobaoWdkEquipmentWcsWcsinfoUploadRequest struct {
    model.Params
    // 上传信息
    param0   string
}

// 初始化TaobaoWdkEquipmentWcsWcsinfoUploadRequest对象
func NewTaobaoWdkEquipmentWcsWcsinfoUploadRequest() *TaobaoWdkEquipmentWcsWcsinfoUploadRequest{
    return &TaobaoWdkEquipmentWcsWcsinfoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentWcsWcsinfoUploadRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.wcs.wcsinfo.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentWcsWcsinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 上传信息
func (r *TaobaoWdkEquipmentWcsWcsinfoUploadRequest) SetParam0(param0 string) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoWdkEquipmentWcsWcsinfoUploadRequest) GetParam0() string {
    return r.param0
}
