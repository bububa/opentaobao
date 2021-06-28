package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星质量分相关接口 APIResponse
taobao.simba.salestar.keywords.qscore.split.get

获取关键词新的质量分
*/
type TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_salestar_keywords_qscore_split_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto `json:"result,omitempty" xml:"