package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询后端商品 APIResponse
taobao.scitem.query

查询后端商品
*/
type TaobaoScitemQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"scitem_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // List<ScItemDO>
    
    ScItemList   []ScItem `json:"sc_item_list,omitempty" xml:"