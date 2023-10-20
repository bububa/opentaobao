package westcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmCustomerInfoGetAPIResponse 会员信息查询接口 API返回值
// alibaba.westcrm.customer.info.get
//
// 会员信息查询接口
type AlibabaWestcrmCustomerInfoGetAPIResponse struct {
	model.CommonResponse
	AlibabaWestcrmCustomerInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWestcrmCustomerInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWestcrmCustomerInfoGetAPIResponseModel).Reset()
}

// AlibabaWestcrmCustomerInfoGetAPIResponseModel is 会员信息查询接口 成功返回结果
type AlibabaWestcrmCustomerInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_customer_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWestcrmCustomerInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWestcrmCustomerInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWestcrmCustomerInfoGetAPIResponse)
	},
}

// GetAlibabaWestcrmCustomerInfoGetAPIResponse 从 sync.Pool 获取 AlibabaWestcrmCustomerInfoGetAPIResponse
func GetAlibabaWestcrmCustomerInfoGetAPIResponse() *AlibabaWestcrmCustomerInfoGetAPIResponse {
	return poolAlibabaWestcrmCustomerInfoGetAPIResponse.Get().(*AlibabaWestcrmCustomerInfoGetAPIResponse)
}

// ReleaseAlibabaWestcrmCustomerInfoGetAPIResponse 将 AlibabaWestcrmCustomerInfoGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWestcrmCustomerInfoGetAPIResponse(v *AlibabaWestcrmCustomerInfoGetAPIResponse) {
	v.Reset()
	poolAlibabaWestcrmCustomerInfoGetAPIResponse.Put(v)
}
