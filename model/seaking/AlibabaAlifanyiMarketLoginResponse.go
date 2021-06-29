package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
登陆用户 APIResponse
alibaba.alifanyi.market.login

企业或组织购买软件服务后可登陆阿里翻译众包系统，接口返回该企业的用户。
*/
type AlibabaAlifanyiMarketLoginAPIResponse struct {
    model.CommonResponse
    AlibabaAlifanyiMarketLoginResponse
}

type AlibabaAlifanyiMarketLoginResponse struct {
    XMLName xml.Name `xml:"alibaba_alifanyi_market_login_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *ResMsgClientDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
