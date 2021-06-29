package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）根据一堆关键词ids获取关键词 API返回值 
taobao.simba.keyword.findbyids

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordFindbyidsAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordFindbyidsResponse
}

// （新）根据一堆关键词ids获取关键词 成功返回结果
type TaobaoSimbaKeywordFindbyidsResponse struct {
    XMLName xml.Name `xml:"simba_keyword_findbyids_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 整体的返回值
    Results   []SiriusBidwordDTO `json:"results,omitempty" xml:"results>sirius_bidword_dto,omitempty"`
    // 错误原因
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
