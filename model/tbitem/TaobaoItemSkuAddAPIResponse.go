package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemskuaddAPIResponse 添加SKU API返回值
// taobao.item.sku.add
//
// 新增一个sku到num_iid指定的商品中 &lt;br/&gt;传入的iid所对应的商品必须属于当前会话的用户
type TaobaoitemskuaddAPIResponse struct {
	model.CommonResponse
	TaobaoitemskuaddAPIResponseModel
}

// TaobaoitemskuaddAPIResponseModel is 添加SKU 成功返回结果
type TaobaoitemskuaddAPIResponseModel struct {
	XMLName xml.Name `xml:"item_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// sku
	Sku *Sku `json:"sku,omitempty" xml:"sku,omitempty"`
}
