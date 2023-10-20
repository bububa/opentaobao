package icbulogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse 四级地址库-街道 API返回值
// alibaba.onetouch.logistics.express.address.street.list
//
// 四级地址库-街道模糊查询
type AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponseModel).Reset()
}

// AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponseModel is 四级地址库-街道 成功返回结果
type AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_address_street_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaOnetouchLogisticsExpressAddressStreetListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse)
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse
func GetAlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse() *AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse {
	return poolAlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse.Get().(*AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse 将 AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse(v *AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse.Put(v)
}
