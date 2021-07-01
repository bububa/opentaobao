package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一个关键词Id列表取得一组关键词 API返回值 
taobao.simba.keywordsbykeywordids.get

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordsbykeywordidsGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordsbykeywordidsGetAPIResponseModel
}

// 根据一个关键词Id列表取得一组关键词 成功返回结果
type TaobaoSimbaKeywordsbykeywordidsGetAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_keywordsbykeywordids_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 取得的关键词列表
    Keywords   []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
}
