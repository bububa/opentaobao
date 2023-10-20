package icbulogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse 四级地址库-市 API返回值
// alibaba.onetouch.logistics.express.address.city.list
//
// 四级地址库-市
type AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressAddressCityListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOnetouchLogisticsExpressAddressCityListAPIResponseModel).Reset()
}

// AlibabaOnetouchLogisticsExpressAddressCityListAPIResponseModel is 四级地址库-市 成功返回结果
type AlibabaOnetouchLogisticsExpressAddressCityListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_address_city_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaOnetouchLogisticsExpressAddressCityListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressAddressCityListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOnetouchLogisticsExpressAddressCityListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse)
	},
}

// GetAlibabaOnetouchLogisticsExpressAddressCityListAPIResponse 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse
func GetAlibabaOnetouchLogisticsExpressAddressCityListAPIResponse() *AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse {
	return poolAlibabaOnetouchLogisticsExpressAddressCityListAPIResponse.Get().(*AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse)
}

// ReleaseAlibabaOnetouchLogisticsExpressAddressCityListAPIResponse 将 AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressAddressCityListAPIResponse(v *AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressAddressCityListAPIResponse.Put(v)
}
