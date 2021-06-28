package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店类目获取接口 APIResponse
taobao.qimen.storecategory.get

商家在ERP中调用该接口，获取门店类目
*/
type TaobaoQimenStorecategoryGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_storecategory_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应信息
    
    Message   string `json:"message,omitempty" xml:"