package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询终端信息 APIResponse
alibaba.aliqin.fc.iot.qrycard

查询终端信息
*/
type AlibabaAliqinFcIotQrycardAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcIotQrycardResponse `json:"alibaba_aliqin_fc_iot_qrycard_response,omitempty"` 
    AlibabaAliqinFcIotQrycardResponse
}

/* model for simplify = false
type AlibabaAliqinFcIotQrycardResponse struct {

    // 系统自动生成
    
    Result  *struct {
        AlibabaAliqinFcIotQrycardResult  *AlibabaAliqinFcIotQrycardResult `json:"alibaba_aliqin_fc_iot_qrycard_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcIotQrycardResponse struct {

    // 系统自动生成
    Result   *AlibabaAliqinFcIotQrycardResult `json:"result,omitempty"`

}
