package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）根据一堆关键词ids获取关键词 APIResponse
taobao.simba.keyword.findbyids

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordFindbyidsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_keyword_findbyids_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 整体的返回值
    
    Results   []SiriusBidwordDto `json:"results,omitempty" xml:"