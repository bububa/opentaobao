package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
货主仓库资源查询接口 APIResponse
taobao.qimen.warehouseinfo.query

货主仓库资源查询
*/
type TaobaoQimenWarehouseinfoQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_warehouseinfo_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"