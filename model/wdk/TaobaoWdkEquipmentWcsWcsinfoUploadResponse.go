package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链业务信息上传 APIResponse
taobao.wdk.equipment.wcs.wcsinfo.upload

五道口仓库悬挂链信息上传
*/
type TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wdk_equipment_wcs_wcsinfo_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // model
    
    Model   string `json:"model,omitempty" xml:"