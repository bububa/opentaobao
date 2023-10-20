package icbulogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse 获取商品类型配置项 API返回值
// alibaba.onetouch.logistics.express.special.product.type.list
//
// 获取商品类型配置项
type AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponseModel).Reset()
}

// AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponseModel is 获取商品类型配置项 成功返回结果
type AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_special_product_type_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse)
	},
}

// GetAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse
func GetAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse() *AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse {
	return poolAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse.Get().(*AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse)
}

// ReleaseAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse 将 AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse(v *AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse.Put(v)
}
