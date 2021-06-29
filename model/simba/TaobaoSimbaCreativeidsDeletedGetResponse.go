package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的创意ID APIResponse
taobao.simba.creativeids.deleted.get

获取删除的创意ID
*/
type TaobaoSimbaCreativeidsDeletedGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCreativeidsDeletedGetResponse
}

type TaobaoSimbaCreativeidsDeletedGetResponse struct {
    XMLName xml.Name `xml:"simba_creativeids_deleted_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 创意ID列表
    
    DeletedCreativeIds   []int64 `json:"deleted_creative_ids,omitempty" xml:"deleted_creative_ids>int64,omitempty"`
    
    
}
