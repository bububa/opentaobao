package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取后台供卖家发布商品的标准商品类目 APIResponse
taobao.itemcats.get

获取后台供卖家发布商品的标准商品类目。
*/
type TaobaoItemcatsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"itemcats_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 最近修改时间(如果取增量，会返回该字段)。
    
    LastModified   string `json:"last_modified,omitempty" xml:"