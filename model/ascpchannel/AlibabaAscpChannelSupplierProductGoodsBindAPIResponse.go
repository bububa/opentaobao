package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsupplierproductgoodsbindAPIResponse 渠道产品与货品绑定接口 API返回值
// alibaba.ascp.channel.supplier.product.goods.bind
//
// 渠道产品与货品绑定接口
type AlibabaascpchannelsupplierproductgoodsbindAPIResponse struct {
	model.CommonResponse
	AlibabaascpchannelsupplierproductgoodsbindAPIResponseModel
}

// AlibabaascpchannelsupplierproductgoodsbindAPIResponseModel is 渠道产品与货品绑定接口 成功返回结果
type AlibabaascpchannelsupplierproductgoodsbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_goods_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
