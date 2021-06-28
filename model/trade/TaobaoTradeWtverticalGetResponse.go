package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
网厅垂直信息查询接口 APIResponse
taobao.trade.wtvertical.get

网厅订单垂直信息的查询
*/
type TaobaoTradeWtverticalGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"trade_wtvertical_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回交易垂直信息的数据结构
    
    WtextResults   []WtExtResult `json:"wtext_results,omitempty" xml:"