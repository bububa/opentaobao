package tmallhk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫国际订单清关信息 API返回值 
tmall.hk.order.clearance.query

天猫国际订单清关信息查询
*/
type TmallHkOrderClearanceQueryAPIResponse struct {
    model.CommonResponse
    TmallHkOrderClearanceQueryResponse
}

// 天猫国际订单清关信息 成功返回结果
type TmallHkOrderClearanceQueryResponse struct {
    XMLName xml.Name `xml:"tmall_hk_order_clearance_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果码
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
    // 结果描述
    ResponseMessage   string `json:"response_message,omitempty" xml:"response_message,omitempty"`
    // 清关数据封装
    Obj   *ClearanceDataDo `json:"obj,omitempty" xml:"obj,omitempty"`
    // 是否正常
    Succeeded   bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}
