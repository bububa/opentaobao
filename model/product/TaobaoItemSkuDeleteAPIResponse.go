package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemSkuDeleteAPIResponse
删除SKU API返回值
taobao.item.sku.delete

删除一个sku的数据<br/>需要删除的sku通过属性properties进行匹配查找 */
type TaobaoItemSkuDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoItemSkuDeleteAPIResponseModel
}

// TaobaoItemSkuDeleteAPIResponseModel is 删除SKU 成功返回结果
type TaobaoItemSkuDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"item_sku_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Sku结构
	Sku *Sku `json:"sku,omitempty" xml:"sku,omitempty"`
}
