package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询采购申请/经销采购单 APIResponse
taobao.fenxiao.dealer.requisitionorder.get

批量查询采购申请/经销采购单，目前支持供应商和分销商查询
*/
type TaobaoFenxiaoDealerRequisitionorderGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_dealer_requisitionorder_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 按查询条件查到的记录总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"