package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除一批关键词 APIResponse
taobao.simba.keywords.delete

删除一批关键词
*/
type TaobaoSimbaKeywordsDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordsDeleteResponse `json:"taobao_simba_keywords_delete_response,omitempty"`
}

type TaobaoSimbaKeywordsDeleteResponse struct {

    // 成功删除的关键词列表
    Keywords   []Keyword `json:"keywords,omitempty"`

}
