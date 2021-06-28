package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫简化发布商品 APIResponse
tmall.item.simpleschema.add

天猫简化版schema发布商品。
*/
type TmallItemSimpleschemaAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_simpleschema_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 发布成功后返回商品ID
    
    Result   string `json:"result,omitempty" xml:"