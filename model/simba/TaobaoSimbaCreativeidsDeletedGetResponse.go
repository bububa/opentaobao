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
	RequestId     string         `json:"request_id,omitempty" xml:"simba_creativeids_deleted_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创意ID列表
    
    DeletedCreativeIds   []int64 `json:"deleted_creative_ids,omitempty" xml:"