package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询物联卡个人实人认证信息 APIResponse
alibaba.aliqin.fc.iot.qry.personinfo

查询物联卡个人实人认证信息
*/
type AlibabaAliqinFcIotQryPersoninfoAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcIotQryPersoninfoResponse `json:"alibaba_aliqin_fc_iot_qry_personinfo_response,omitempty"` 
    AlibabaAliqinFcIotQryPersoninfoResponse
}

/* model for simplify = false
type AlibabaAliqinFcIotQryPersoninfoResponse struct {

    // result
    
    Result  *struct {
        AlibabaAliqinFcIotQryPersoninfoResult  *AlibabaAliqinFcIotQryPersoninfoResult `json:"alibaba_aliqin_fc_iot_qry_personinfo_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcIotQryPersoninfoResponse struct {

    // result
    Result   *AlibabaAliqinFcIotQryPersoninfoResult `json:"result,omitempty"`

}
