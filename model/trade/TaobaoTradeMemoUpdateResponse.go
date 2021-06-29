package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改交易备注 APIResponse
taobao.trade.memo.update

需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能
*/
type TaobaoTradeMemoUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoTradeMemoUpdateResponse
}

type TaobaoTradeMemoUpdateResponse struct {
    XMLName xml.Name `xml:"trade_memo_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新交易的备注信息后返回的Trade，其中可用字段为tid和modified
    
    Trade   *Trade `json:"trade,omitempty" xml:"trade,omitempty"`

    
}
