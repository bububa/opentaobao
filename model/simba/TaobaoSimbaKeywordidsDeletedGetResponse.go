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
    Response *TaobaoSimbaKeywordidsDeletedGetResponse `json:"taobao_simba_keywordids_deleted_get_response,omitempty"`
}

type TaobaoSimbaKeywordidsDeletedGetResponse struct {

    // 词ID列表
    DeletedKeywordIds   []Number `json:"deleted_keyword_ids,omitempty"`

}
