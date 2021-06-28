package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓库同步接口 APIResponse
taobao.qimen.warehouseinfo.synchronize

仓库同步接口
*/
type TaobaoQimenWarehouseinfoSynchronizeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_warehouseinfo_synchronize_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应报文
    
    Response   *Response `json:"response,omitempty" xml:"