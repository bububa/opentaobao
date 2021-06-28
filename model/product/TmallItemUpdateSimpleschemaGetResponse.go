package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
官网同购编辑商品的get接口 APIResponse
tmall.item.update.simpleschema.get

官网同购编辑商品的get接口
*/
type TmallItemUpdateSimpleschemaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_update_simpleschema_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Error   bool `json:"error,omitempty" xml:"