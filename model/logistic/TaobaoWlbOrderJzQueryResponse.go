package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
家装业务查询物流公司api APIResponse
taobao.wlb.order.jz.query

家装业务查询物流公司api
*/
type TaobaoWlbOrderJzQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_order_jz_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误编码
    
    ResultErrorCode   string `json:"result_error_code,omitempty" xml:"