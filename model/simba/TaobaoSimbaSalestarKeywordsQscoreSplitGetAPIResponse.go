package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星质量分相关接口 API返回值 
taobao.simba.salestar.keywords.qscore.split.get

获取关键词新的质量分
*/
type TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponseModel
}

// (新)销量明星质量分相关接口 成功返回结果
type TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_salestar_keywords_qscore_split_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
