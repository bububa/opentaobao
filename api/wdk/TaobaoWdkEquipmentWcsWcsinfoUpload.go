package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
悬挂链业务信息上传 
taobao.wdk.equipment.wcs.wcsinfo.upload

五道口仓库悬挂链信息上传
*/
func TaobaoWdkEquipmentWcsWcsinfoUpload(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentWcsWcsinfoUploadRequest, session string) (*wdk.TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse, error) {
    var resp wdk.TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
