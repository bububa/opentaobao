package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询活动 API返回值 
alibaba.lsy.crm.activity.getbaseinfo

ISV查询活动
*/
type AlibabaLsyCrmActivityGetbaseinfoAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmActivityGetbaseinfoResponse
}

// ISV查询活动 成功返回结果
type AlibabaLsyCrmActivityGetbaseinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_activity_getbaseinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *ResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
