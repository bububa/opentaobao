package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemskuupdateAPIResponse 更新SKU信息 API返回值
// taobao.item.sku.update
//
// *更新一个sku的数据 &lt;br/&gt;*需要更新的sku通过属性properties进行匹配查找 &lt;br/&gt;*商品的数量和价格必须大于等于0 &lt;br/&gt;*sku记录会更新到指定的num_iid对应的商品中 &lt;br/&gt;*num_iid对应的商品必须属于当前的会话用户
type TaobaoitemskuupdateAPIResponse struct {
	model.CommonResponse
	TaobaoitemskuupdateAPIResponseModel
}

// TaobaoitemskuupdateAPIResponseModel is 更新SKU信息 成功返回结果
type TaobaoitemskuupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"item_sku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品Sku
	Sku *Sku `json:"sku,omitempty" xml:"sku,omitempty"`
}
