package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）关键词更新相关接口 API返回值 
taobao.simba.keyword.update

（新）关键词更新相关接口
*/
type TaobaoSimbaKeywordUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordUpdateResponse
}

// （新）关键词更新相关接口 成功返回结果
type TaobaoSimbaKeywordUpdateResponse struct {
    XMLName xml.Name `xml:"simba_keyword_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 整体的返回值
    Results   []SiriusBidwordDto `json:"results,omitempty" xml:"results>sirius_bidword_dto,omitempty"`
    // 错误原因
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
