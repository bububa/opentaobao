package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商修改申请单 API返回值 
tmall.channel.trade.applyorder.modify

上游供应商修改申请单, 目前只允许修改价格+件数且sku数量必须完全一致
*/
type TmallChannelTradeApplyorderModifyAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeApplyorderModifyAPIResponseModel
}

// 供应商修改申请单 成功返回结果
type TmallChannelTradeApplyorderModifyAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_channel_trade_applyorder_modify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallChannelTradeApplyorderModifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
