package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物联卡个人实人认证信息 APIResponse
alibaba.aliqin.fc.iot.qry.personinfo

查询物联卡个人实人认证信息
*/
type AlibabaAliqinFcIotQryPersoninfoAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotQryPersoninfoResponse
}

type AlibabaAliqinFcIotQryPersoninfoResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_qry_personinfo_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaAliqinFcIotQryPersoninfoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
