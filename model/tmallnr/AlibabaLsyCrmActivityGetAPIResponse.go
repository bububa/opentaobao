package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购查询活动详情 API返回值 
alibaba.lsy.crm.activity.get

私域导购查询活动详情
*/
type AlibabaLsyCrmActivityGetAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmActivityGetAPIResponseModel
}

// 私域导购查询活动详情 成功返回结果
type AlibabaLsyCrmActivityGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_activity_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值
    Result   *AlibabaLsyCrmActivityGetResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
