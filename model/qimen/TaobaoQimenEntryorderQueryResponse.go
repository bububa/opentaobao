package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
入库单查询接口 APIResponse
taobao.qimen.entryorder.query

ERP调用接口，查询入库单信息;
*/
type TaobaoQimenEntryorderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_entryorder_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *EntryOrderQueryResponse `json:"response,omitempty" xml:"