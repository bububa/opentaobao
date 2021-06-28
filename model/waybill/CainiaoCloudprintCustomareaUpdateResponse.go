package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自定义区内容更新 APIResponse
cainiao.cloudprint.customarea.update

自定义区内容更新
*/
type CainiaoCloudprintCustomareaUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cloudprint_customarea_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"