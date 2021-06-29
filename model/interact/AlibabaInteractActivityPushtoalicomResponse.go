package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
小铺isv推广流量活动到流量聚乐部 API返回值 
alibaba.interact.activity.pushtoalicom

涉及到流量包的小铺isv，将活动推送到流量聚乐部
*/
type AlibabaInteractActivityPushtoalicomAPIResponse struct {
    model.CommonResponse
    AlibabaInteractActivityPushtoalicomResponse
}

// 小铺isv推广流量活动到流量聚乐部 成功返回结果
type AlibabaInteractActivityPushtoalicomResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_activity_pushtoalicom_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 推送成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
