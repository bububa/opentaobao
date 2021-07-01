package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查面单号状态v1.0 API返回值 
taobao.wlb.waybill.i.querydetail

查看面单号的当前状态，如签收、发货、失效等。
*/
type TaobaoWlbWaybillIQuerydetailAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWaybillIQuerydetailAPIResponseModel
}

// 查面单号状态v1.0 成功返回结果
type TaobaoWlbWaybillIQuerydetailAPIResponseModel struct {
    XMLName xml.Name `xml:"wlb_waybill_i_querydetail_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 不存在的面单号
    InexistentWaybillCodes   []string `json:"inexistent_waybill_codes,omitempty" xml:"inexistent_waybill_codes>string,omitempty"`
    // 查询是否成功
    QuerySuccess   bool `json:"query_success,omitempty" xml:"query_success,omitempty"`
    // 面单详情
    WaybillDetails   []WaybillDetailQueryInfo `json:"waybill_details,omitempty" xml:"waybill_details>waybill_detail_query_info,omitempty"`
    // 面单查询错误编码
    ErrorCodes   []string `json:"error_codes,omitempty" xml:"error_codes>string,omitempty"`
}
