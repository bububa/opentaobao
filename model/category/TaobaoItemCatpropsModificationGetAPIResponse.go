package category

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemcatpropsmodificationgetAPIResponse 查询商品类目属性变更 API返回值
// taobao.item.catprops.modification.get
//
// 查询商品类目属性变更信息
type TaobaoitemcatpropsmodificationgetAPIResponse struct {
	model.CommonResponse
	TaobaoitemcatpropsmodificationgetAPIResponseModel
}

// TaobaoitemcatpropsmodificationgetAPIResponseModel is 查询商品类目属性变更 成功返回结果
type TaobaoitemcatpropsmodificationgetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_catprops_modification_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Results []PropsModificationResult `json:"results,omitempty" xml:"results>props_modification_result,omitempty"`
}
