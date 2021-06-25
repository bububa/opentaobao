package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建一批关键词 APIResponse
taobao.simba.keywordsvon.add

创建一批关键词
*/
type TaobaoSimbaKeywordsvonAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordsvonAddResponse `json:"taobao_simba_keywordsvon_add_response,omitempty"`
}

type TaobaoSimbaKeywordsvonAddResponse struct {

    // 关键词列表
    Keywords   []Keyword `json:"keywords,omitempty"`

}
