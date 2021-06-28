package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
失效指定用户的商品与后端商品的映射关系 APIResponse
taobao.scitem.map.delete

根据后端商品Id，失效指定用户的商品与后端商品的映射关系
*/
type TaobaoScitemMapDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"scitem_map_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 失效条数
    
    Module   int64 `json:"module,omitempty" xml:"