package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDistributorProductQuantityGetAPIResponse 分销商查询产品库存 API返回值
// taobao.fenxiao.distributor.product.quantity.get
//
// 分销商查询产品库存
type TaobaoFenxiaoDistributorProductQuantityGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDistributorProductQuantityGetAPIResponseModel
}

// TaobaoFenxiaoDistributorProductQuantityGetAPIResponseModel is 分销商查询产品库存 成功返回结果
type TaobaoFenxiaoDistributorProductQuantityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_distributor_product_quantity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
