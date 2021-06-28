package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的词ID APIResponse
taobao.simba.keywordids.deleted.get

获取删除的词ID
*/
type TaobaoSimbaKeywordidsDeletedGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaKeywordidsDeletedGetResponse `json:"simba_keywordids_deleted_get_response,omitempty"` 
    TaobaoSimbaKeywordidsDeletedGetResponse
}

/* model for simplify = false
type TaobaoSimbaKeywordidsDeletedGetResponse struct {

    // 词ID列表
    
    DeletedKeywordIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"deleted_keyword_ids,omitempty"`
    

}
*/

type TaobaoSimbaKeywordidsDeletedGetResponse struct {

    // 词ID列表
    DeletedKeywordIds   []int64 `json:"deleted_keyword_ids,omitempty"`

}
