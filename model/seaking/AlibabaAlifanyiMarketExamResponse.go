package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过考试用户 APIResponse
alibaba.alifanyi.market.exam

企业或组织购买软件服务后可参与阿里翻译在线系统的考试认证，接口返回该企业或组织认证通过的用户
*/
type AlibabaAlifanyiMarketExamAPIResponse struct {
    model.CommonResponse
    AlibabaAlifanyiMarketExamResponse
}

type AlibabaAlifanyiMarketExamResponse struct {
    XMLName xml.Name `xml:"alibaba_alifanyi_market_exam_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *ResMsgClientDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
