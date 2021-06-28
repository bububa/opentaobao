package yunos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
YUNOS生活服务群发消息 APIResponse
yunos.cloudcard.batch.opermsg.send

这个是一个群发消息接口，ISV通过该接口给订阅自己服务号的所有YUNOS终端用户发送服务号消息，目前该接口有调用频率限制，具体规则参考YUNOS开放平台文档。
*/
type YunosCloudcardBatchOpermsgSendAPIResponse struct {
    model.CommonResponse
    YunosCloudcardBatchOpermsgSendResponse
}

type YunosCloudcardBatchOpermsgSendResponse struct {
    XMLName xml.Name `xml:"yunos_cloudcard_batch_opermsg_send_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 群发消息是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
