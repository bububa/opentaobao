package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeyworddeleteAPIResponse 外贸直通车删除关键词 API返回值
// alibaba.scbp.ad.keyword.delete
//
// 外贸直通车删除关键词
type AlibabascbpadkeyworddeleteAPIResponse struct {
	model.CommonResponse
	AlibabascbpadkeyworddeleteAPIResponseModel
}

// AlibabascbpadkeyworddeleteAPIResponseModel is 外贸直通车删除关键词 成功返回结果
type AlibabascbpadkeyworddeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除关键词是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
