package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词所属分组 API返回值 
alibaba.scbp.ad.keyword.tag.update

修改关键词所属分组
*/
type AlibabaScbpAdKeywordTagUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordTagUpdateAPIResponseModel
}

// 修改关键词所属分组 成功返回结果
type AlibabaScbpAdKeywordTagUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_tag_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 实际修改的关键词数
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
}
