package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
登陆用户 API返回值 
alibaba.alifanyi.market.login

企业或组织购买软件服务后可登陆阿里翻译众包系统，接口返回该企业的用户。
*/
type AlibabaAlifanyiMarketLoginAPIResponse struct {
    model.CommonResponse
    AlibabaAlifanyiMarketLoginResponse
}

// 登陆用户 成功返回结果
type AlibabaAlifanyiMarketLoginResponse struct {
    XMLName xml.Name `xml:"alibaba_alifanyi_market_login_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *ResMsgClientDto `json:"result,omitempty" xml:"result,omitempty"`
}
