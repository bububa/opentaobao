package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单基本信息 APIResponse
alibaba.ele.fengniao.order.query

查询订单基本信息
*/
type AlibabaEleFengniaoOrderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ele_fengniao_order_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // kvs
    
    Kvs   []Kvs `json:"kvs,omitempty" xml:"