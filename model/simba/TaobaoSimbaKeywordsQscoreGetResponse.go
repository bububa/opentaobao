package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分 APIResponse
taobao.simba.keywords.qscore.get

取得一个推广组的所有关键词的质量得分列表
*/
type TaobaoSimbaKeywordsQscoreGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordsQscoreGetResponse
}

type TaobaoSimbaKeywordsQscoreGetResponse struct {
    XMLName xml.Name `xml:"simba_keywords_qscore_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 取得的关键词质量得分列表
    
    KeywordQscoreList   []KeywordQscore `json:"keyword_qscore_list,omitempty" xml:"keyword_qscore_list>keyword_qscore,omitempty"`
    
    
}
