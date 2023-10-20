package category

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaImapFixedmappingQueryAPIResponse 查询两个渠道之间的固定映射关系，不通过算法兜底 API返回值
// alibaba.imap.fixedmapping.query
//
// 查询两个渠道之间的固定映射关系，不通过算法兜底
type AlibabaImapFixedmappingQueryAPIResponse struct {
	model.CommonResponse
	AlibabaImapFixedmappingQueryAPIResponseModel
}

// AlibabaImapFixedmappingQueryAPIResponseModel is 查询两个渠道之间的固定映射关系，不通过算法兜底 成功返回结果
type AlibabaImapFixedmappingQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_imap_fixedmapping_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaImapFixedmappingQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
