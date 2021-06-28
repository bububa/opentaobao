package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的词ID APIResponse
taobao.simba.keywordids.deleted.get

获取删除的词ID
*/
type TaobaoSimbaKeywordidsDeletedGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_keywordids_deleted_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 词ID列表
    
    DeletedKeywordIds   []int64 `json:"deleted_keyword_ids,omitempty" xml:"