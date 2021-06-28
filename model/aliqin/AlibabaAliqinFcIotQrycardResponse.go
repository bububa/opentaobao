package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询终端信息 APIResponse
alibaba.aliqin.fc.iot.qrycard

查询终端信息
*/
type AlibabaAliqinFcIotQrycardAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotQrycardResponse
}

type AlibabaAliqinFcIotQrycardResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_qrycard_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Result   *AlibabaAliqinFcIotQrycardResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
