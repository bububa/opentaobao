package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取服务商品扩展信息 APIResponse
tmall.fuwu.serviceitem.list

获取服务商品扩展信息
*/
type TmallFuwuServiceitemListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_fuwu_serviceitem_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallFuwuServiceitemListResult `json:"result,omitempty" xml:"