package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 APIResponse
taobao.qimen.singleitem.synchronize

ERP调用奇门的接口,同步商品信息给WMS
*/
type TaobaoQimenSingleitemSynchronizeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_singleitem_synchronize_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"