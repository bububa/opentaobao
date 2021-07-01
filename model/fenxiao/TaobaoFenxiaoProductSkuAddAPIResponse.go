package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductSkuAddAPIResponse
产品sku添加接口 API返回值
taobao.fenxiao.product.sku.add

添加产品SKU信息 */
type TaobaoFenxiaoProductSkuAddAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductSkuAddAPIResponseModel
}

// TaobaoFenxiaoProductSkuAddAPIResponseModel is 产品sku添加接口 成功返回结果
type TaobaoFenxiaoProductSkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 操作时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
}
