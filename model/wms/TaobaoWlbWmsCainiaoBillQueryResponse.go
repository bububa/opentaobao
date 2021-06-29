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
    TaobaoWlbWmsCainiaoBillQueryResponse
}

type TaobaoWlbWmsCainiaoBillQueryResponse struct {
    XMLName xml.Name `xml:"wlb_wms_cainiao_bill_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 总条数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
    // 订单列表信息
    
    OrderInfoList   []CainiaoBillQueryOrderinfolist `json:"order_info_list,omitempty" xml:"order_info_list>cainiao_bill_query_orderinfolist,omitempty"`
    
    
}
