package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取修改的词ID APIResponse
taobao.simba.keywordids.changed.get

获取修改的词ID
*/
type TaobaoSimbaKeywordidsChangedGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaKeywordidsChangedGetResponse `json:"simba_keywordids_changed_get_response,omitempty"` 
    TaobaoSimbaKeywordidsChangedGetResponse
}

/* model for simplify = false
type TaobaoSimbaKeywordidsChangedGetResponse struct {

    // 词的ID列表
    
    ChangedKeywordIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"changed_keyword_ids,omitempty"`
    

}
*/

type TaobaoSimbaKeywordidsChangedGetResponse struct {

    // 词的ID列表
    ChangedKeywordIds   []int64 `json:"changed_keyword_ids,omitempty"`

}
