package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpMaterialItemFindpageAPIResponse 分页查询商品信息 API返回值
// taobao.universalbp.material.item.findpage
//
// 分页获取店铺内的商品列表
type TaobaoUniversalbpMaterialItemFindpageAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpMaterialItemFindpageAPIResponseModel
}

// TaobaoUniversalbpMaterialItemFindpageAPIResponseModel is 分页查询商品信息 成功返回结果
type TaobaoUniversalbpMaterialItemFindpageAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_material_item_findpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpMaterialItemFindpageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
