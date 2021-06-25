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
    Response *TaobaoSimbaKeywordidsChangedGetResponse `json:"taobao_simba_keywordids_changed_get_response,omitempty"`
}

type TaobaoSimbaKeywordidsChangedGetResponse struct {

    // 词的ID列表
    ChangedKeywordIds   []Number `json:"changed_keyword_ids,omitempty"`

}
