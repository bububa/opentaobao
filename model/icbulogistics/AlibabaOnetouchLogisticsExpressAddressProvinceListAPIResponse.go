package icbulogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse 四级地址库-省 API返回值
// alibaba.onetouch.logistics.express.address.province.list
//
// 四级地址库-省
type AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponseModel).Reset()
}

// AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponseModel is 四级地址库-省 成功返回结果
type AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_address_province_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaOnetouchLogisticsExpressAddressProvinceListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse)
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse
func GetAlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse() *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse {
	return poolAlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse.Get().(*AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse 将 AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse(v *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse.Put(v)
}
