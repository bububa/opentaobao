package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单据列表 APIResponse
taobao.wlb.wms.cainiao.bill.query

查询单据列表
*/
type TaobaoWlbWmsCainiaoBillQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_cainiao_bill_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 总条数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"