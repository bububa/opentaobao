package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取信息流账户详情 APIResponse
taobao.feedflow.account.get

获取账户信息接口。
(1) BP显示余额 (字段 ：banlance ) = 现金余额(字段：cash_balance) + 赠款余额；
(2) 可用余额(字段：availableBalance) = BP显示余额
(3) 红包(字段：redPacket)
*/
type TaobaoFeedflowAccountGetAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowAccountGetResponse
}

type TaobaoFeedflowAccountGetResponse struct {
    XMLName xml.Name `xml:"feedflow_account_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`

    
}
