package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
插入和更新食谱 APIResponse
alibaba.ailabs.iot.business.recipe.insertorupdate

插入和更新食谱，将isv的食谱添加到云端进行存储
*/
type AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ailabs_iot_business_recipe_insertorupdate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 信息
    
    Message   string `json:"message,omitempty" xml:"