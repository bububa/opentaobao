package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词和类目出价的质量得分 APIResponse
taobao.simba.keywordscat.qscore.get

取得一个推广组的所有关键词和类目出价的质量得分列表
*/
type TaobaoSimbaKeywordscatQscoreGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_keywordscat_qscore_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 类目出价和词的质量得分对象
    
    Qscore   *Qscore `json:"qscore,omitempty" xml:"