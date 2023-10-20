package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse 删除屏蔽品 API返回值
// alibaba.scbp.ad.group.delete.forbidden.product
//
// 删除屏蔽品
type AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdGroupDeleteForbiddenProductAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdGroupDeleteForbiddenProductAPIResponseModel).Reset()
}

// AlibabaScbpAdGroupDeleteForbiddenProductAPIResponseModel is 删除屏蔽品 成功返回结果
type AlibabaScbpAdGroupDeleteForbiddenProductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_delete_forbidden_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupDeleteForbiddenProductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdGroupDeleteForbiddenProductAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse)
	},
}

// GetAlibabaScbpAdGroupDeleteForbiddenProductAPIResponse 从 sync.Pool 获取 AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse
func GetAlibabaScbpAdGroupDeleteForbiddenProductAPIResponse() *AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse {
	return poolAlibabaScbpAdGroupDeleteForbiddenProductAPIResponse.Get().(*AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse)
}

// ReleaseAlibabaScbpAdGroupDeleteForbiddenProductAPIResponse 将 AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdGroupDeleteForbiddenProductAPIResponse(v *AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdGroupDeleteForbiddenProductAPIResponse.Put(v)
}
