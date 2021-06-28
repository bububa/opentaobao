package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
入库单创建接口 APIResponse
taobao.qimen.entryorder.create

ERP调用接口，创建入库单;
*/
type TaobaoQimenEntryorderCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_entryorder_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"