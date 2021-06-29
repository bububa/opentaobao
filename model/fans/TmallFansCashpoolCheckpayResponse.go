package fans

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
检查资金池付款状态 API返回值 
tmall.fans.cashpool.checkpay

检查资金池付款状态
*/
type TmallFansCashpoolCheckpayAPIResponse struct {
    model.CommonResponse
    TmallFansCashpoolCheckpayResponse
}

// 检查资金池付款状态 成功返回结果
type TmallFansCashpoolCheckpayResponse struct {
    XMLName xml.Name `xml:"tmall_fans_cashpool_checkpay_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值
    FansResult   *FansResult `json:"fans_result,omitempty" xml:"fans_result,omitempty"`
}
