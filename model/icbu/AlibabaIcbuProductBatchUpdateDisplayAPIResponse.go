package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductBatchUpdateDisplayAPIResponse 商品批量上下架接口 API返回值
// alibaba.icbu.product.batch.update.display
//
// 给国际站的三方服务商提供批量上下架接口
type AlibabaIcbuProductBatchUpdateDisplayAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductBatchUpdateDisplayAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuProductBatchUpdateDisplayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuProductBatchUpdateDisplayAPIResponseModel).Reset()
}

// AlibabaIcbuProductBatchUpdateDisplayAPIResponseModel is 商品批量上下架接口 成功返回结果
type AlibabaIcbuProductBatchUpdateDisplayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_batch_update_display_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 只有出错才会显示，唯一标识这次请求
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 如果出错，这里会显示错误码
	SubErrorCode string `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
	// 具体出错信息
	SubErrorMsg string `json:"sub_error_msg,omitempty" xml:"sub_error_msg,omitempty"`
	// 本次操作是否成功，true表示成功，false表示失败
	SubSuccess bool `json:"sub_success,omitempty" xml:"sub_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuProductBatchUpdateDisplayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.SubErrorCode = ""
	m.SubErrorMsg = ""
	m.SubSuccess = false
}

var poolAlibabaIcbuProductBatchUpdateDisplayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuProductBatchUpdateDisplayAPIResponse)
	},
}

// GetAlibabaIcbuProductBatchUpdateDisplayAPIResponse 从 sync.Pool 获取 AlibabaIcbuProductBatchUpdateDisplayAPIResponse
func GetAlibabaIcbuProductBatchUpdateDisplayAPIResponse() *AlibabaIcbuProductBatchUpdateDisplayAPIResponse {
	return poolAlibabaIcbuProductBatchUpdateDisplayAPIResponse.Get().(*AlibabaIcbuProductBatchUpdateDisplayAPIResponse)
}

// ReleaseAlibabaIcbuProductBatchUpdateDisplayAPIResponse 将 AlibabaIcbuProductBatchUpdateDisplayAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuProductBatchUpdateDisplayAPIResponse(v *AlibabaIcbuProductBatchUpdateDisplayAPIResponse) {
	v.Reset()
	poolAlibabaIcbuProductBatchUpdateDisplayAPIResponse.Put(v)
}
