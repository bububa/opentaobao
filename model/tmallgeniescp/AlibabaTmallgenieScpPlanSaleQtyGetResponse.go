package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
12-同步销售数据 APIResponse
alibaba.tmallgenie.scp.plan.sale.qty.get

同步销售数据
*/
type AlibabaTmallgenieScpPlanSaleQtyGetAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanSaleQtyGetResponse
}

type AlibabaTmallgenieScpPlanSaleQtyGetResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_sale_qty_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
