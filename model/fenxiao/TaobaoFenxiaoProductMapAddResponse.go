package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建分销和后端商品映射关系 APIResponse
taobao.fenxiao.product.map.add

创建分销和供应链商品映射关系。
*/
type TaobaoFenxiaoProductMapAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_product_map_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"