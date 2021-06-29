package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新质量分服务 APIResponse
taobao.simba.keywords.qscore.split.get

获取关键词新的质量分
*/
type TaobaoSimbaKeywordsQscoreSplitGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordsQscoreSplitGetResponse
}

type TaobaoSimbaKeywordsQscoreSplitGetResponse struct {
    XMLName xml.Name `xml:"simba_keywords_qscore_split_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoSimbaKeywordsQscoreSplitGetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
