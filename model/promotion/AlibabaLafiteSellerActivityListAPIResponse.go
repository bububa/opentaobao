package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLafiteSellerActivityListAPIResponse 商家自运营活动列表 API返回值
// alibaba.lafite.seller.activity.list
//
// 商家查询自己配置的活动列表
type AlibabaLafiteSellerActivityListAPIResponse struct {
	model.CommonResponse
	AlibabaLafiteSellerActivityListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLafiteSellerActivityListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLafiteSellerActivityListAPIResponseModel).Reset()
}

// AlibabaLafiteSellerActivityListAPIResponseModel is 商家自运营活动列表 成功返回结果
type AlibabaLafiteSellerActivityListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lafite_seller_activity_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *AlibabaLafiteSellerActivityListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLafiteSellerActivityListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLafiteSellerActivityListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLafiteSellerActivityListAPIResponse)
	},
}

// GetAlibabaLafiteSellerActivityListAPIResponse 从 sync.Pool 获取 AlibabaLafiteSellerActivityListAPIResponse
func GetAlibabaLafiteSellerActivityListAPIResponse() *AlibabaLafiteSellerActivityListAPIResponse {
	return poolAlibabaLafiteSellerActivityListAPIResponse.Get().(*AlibabaLafiteSellerActivityListAPIResponse)
}

// ReleaseAlibabaLafiteSellerActivityListAPIResponse 将 AlibabaLafiteSellerActivityListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLafiteSellerActivityListAPIResponse(v *AlibabaLafiteSellerActivityListAPIResponse) {
	v.Reset()
	poolAlibabaLafiteSellerActivityListAPIResponse.Put(v)
}
