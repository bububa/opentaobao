package icbulogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse
查询物流运力列表 API返回值
alibaba.onetouch.logistics.express.logistics.product.list

查询物流产品&揽收仓库列表 */
type AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponseModel
}

// AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponseModel is 查询物流运力列表 成功返回结果
type AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_logistics_product_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaOnetouchLogisticsExpressLogisticsProductListResult `json:"result,omitempty" xml:"result,omitempty"`
}
