package icbulogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse 四级地址库-区域 API返回值
// alibaba.onetouch.logistics.express.address.division.list
//
// 四级地址库-区
type AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponseModel).Reset()
}

// AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponseModel is 四级地址库-区域 成功返回结果
type AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_address_division_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaOnetouchLogisticsExpressAddressDivisionListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse)
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse
func GetAlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse() *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse {
	return poolAlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse.Get().(*AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse 将 AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse(v *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse.Put(v)
}
