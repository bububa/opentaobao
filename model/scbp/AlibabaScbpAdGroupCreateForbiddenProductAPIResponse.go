package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupCreateForbiddenProductAPIResponse 创建屏蔽品 API返回值
// alibaba.scbp.ad.group.create.forbidden.product
//
// 创建屏蔽品
type AlibabaScbpAdGroupCreateForbiddenProductAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdGroupCreateForbiddenProductAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupCreateForbiddenProductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdGroupCreateForbiddenProductAPIResponseModel).Reset()
}

// AlibabaScbpAdGroupCreateForbiddenProductAPIResponseModel is 创建屏蔽品 成功返回结果
type AlibabaScbpAdGroupCreateForbiddenProductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_create_forbidden_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupCreateForbiddenProductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdGroupCreateForbiddenProductAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdGroupCreateForbiddenProductAPIResponse)
	},
}

// GetAlibabaScbpAdGroupCreateForbiddenProductAPIResponse 从 sync.Pool 获取 AlibabaScbpAdGroupCreateForbiddenProductAPIResponse
func GetAlibabaScbpAdGroupCreateForbiddenProductAPIResponse() *AlibabaScbpAdGroupCreateForbiddenProductAPIResponse {
	return poolAlibabaScbpAdGroupCreateForbiddenProductAPIResponse.Get().(*AlibabaScbpAdGroupCreateForbiddenProductAPIResponse)
}

// ReleaseAlibabaScbpAdGroupCreateForbiddenProductAPIResponse 将 AlibabaScbpAdGroupCreateForbiddenProductAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdGroupCreateForbiddenProductAPIResponse(v *AlibabaScbpAdGroupCreateForbiddenProductAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdGroupCreateForbiddenProductAPIResponse.Put(v)
}
