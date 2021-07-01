package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商审核拒绝采购申请单 API返回值 
tmall.channel.trade.applyorder.refuse

供应商审核拒绝采购申请单
*/
type TmallChannelTradeApplyorderRefuseAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeApplyorderRefuseAPIResponseModel
}

// 供应商审核拒绝采购申请单 成功返回结果
type TmallChannelTradeApplyorderRefuseAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_channel_trade_applyorder_refuse_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallChannelTradeApplyorderRefuseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
