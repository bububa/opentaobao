package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWholesaleGoodsGetAPIResponse 查询阿里巴巴批发市场商品详情 API返回值
// alibaba.wholesale.goods.get
//
// 查询阿里巴巴批发市场商品详情
type AlibabaWholesaleGoodsGetAPIResponse struct {
	model.CommonResponse
	AlibabaWholesaleGoodsGetAPIResponseModel
}

// AlibabaWholesaleGoodsGetAPIResponseModel is 查询阿里巴巴批发市场商品详情 成功返回结果
type AlibabaWholesaleGoodsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wholesale_goods_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// wholesale goods detail result
	WholesaleGoodsResult *WholesaleGoodsOpenResult `json:"wholesale_goods_result,omitempty" xml:"wholesale_goods_result,omitempty"`
}
