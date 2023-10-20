package icbulogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse 查询物流运力列表 API返回值
// alibaba.onetouch.logistics.express.logistics.product.list
//
// 查询物流产品&amp;揽收仓库列表
type AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponseModel).Reset()
}

// AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponseModel is 查询物流运力列表 成功返回结果
type AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_logistics_product_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaOnetouchLogisticsExpressLogisticsProductListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse)
	},
}

// GetAlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse
func GetAlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse() *AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse {
	return poolAlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse.Get().(*AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse)
}

// ReleaseAlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse 将 AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse(v *AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse.Put(v)
}
