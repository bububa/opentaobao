package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductskusgetAPIResponse SKU查询接口 API返回值
// taobao.fenxiao.product.skus.get
//
// 产品sku查询
type TaobaofenxiaoproductskusgetAPIResponse struct {
	model.CommonResponse
	TaobaofenxiaoproductskusgetAPIResponseModel
}

// TaobaofenxiaoproductskusgetAPIResponseModel is SKU查询接口 成功返回结果
type TaobaofenxiaoproductskusgetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_skus_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// sku信息
	Skus []FenxiaoSku `json:"skus,omitempty" xml:"skus>fenxiao_sku,omitempty"`
	// 记录数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
