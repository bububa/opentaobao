package icbulogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse
四级地址库-省 API返回值
alibaba.onetouch.logistics.express.address.province.list

四级地址库-省 */
type AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponseModel
}

// AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponseModel is 四级地址库-省 成功返回结果
type AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_address_province_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaOnetouchLogisticsExpressAddressProvinceListResult `json:"result,omitempty" xml:"result,omitempty"`
}
