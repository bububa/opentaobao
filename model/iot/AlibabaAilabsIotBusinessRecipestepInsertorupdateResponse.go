package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
插入或更新食谱步骤 APIResponse
alibaba.ailabs.iot.business.recipestep.insertorupdate

插入或更新食谱步骤
*/
type AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ailabs_iot_business_recipestep_insertorupdate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 信息
    
    Message   string `json:"message,omitempty" xml:"