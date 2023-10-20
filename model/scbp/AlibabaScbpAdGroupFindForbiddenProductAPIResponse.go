package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupFindForbiddenProductAPIResponse 查询屏蔽品 API返回值
// alibaba.scbp.ad.group.find.forbidden.product
//
// 查询屏蔽品
type AlibabaScbpAdGroupFindForbiddenProductAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdGroupFindForbiddenProductAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupFindForbiddenProductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdGroupFindForbiddenProductAPIResponseModel).Reset()
}

// AlibabaScbpAdGroupFindForbiddenProductAPIResponseModel is 查询屏蔽品 成功返回结果
type AlibabaScbpAdGroupFindForbiddenProductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_find_forbidden_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回列表
	ResultList []ForbiddenProductDto `json:"result_list,omitempty" xml:"result_list>forbidden_product_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupFindForbiddenProductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpAdGroupFindForbiddenProductAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdGroupFindForbiddenProductAPIResponse)
	},
}

// GetAlibabaScbpAdGroupFindForbiddenProductAPIResponse 从 sync.Pool 获取 AlibabaScbpAdGroupFindForbiddenProductAPIResponse
func GetAlibabaScbpAdGroupFindForbiddenProductAPIResponse() *AlibabaScbpAdGroupFindForbiddenProductAPIResponse {
	return poolAlibabaScbpAdGroupFindForbiddenProductAPIResponse.Get().(*AlibabaScbpAdGroupFindForbiddenProductAPIResponse)
}

// ReleaseAlibabaScbpAdGroupFindForbiddenProductAPIResponse 将 AlibabaScbpAdGroupFindForbiddenProductAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdGroupFindForbiddenProductAPIResponse(v *AlibabaScbpAdGroupFindForbiddenProductAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdGroupFindForbiddenProductAPIResponse.Put(v)
}
