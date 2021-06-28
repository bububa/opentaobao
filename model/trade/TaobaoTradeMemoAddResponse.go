package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对一笔交易添加备注 APIResponse
taobao.trade.memo.add

根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
*/
type TaobaoTradeMemoAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"trade_memo_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 对一笔交易添加备注后返回其对应的Trade，Trade中可用的返回字段有tid和created
    
    Trade   *Trade `json:"trade,omitempty" xml:"