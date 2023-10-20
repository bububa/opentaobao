package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemdeleteAPIResponse 删除单条商品 API返回值
// taobao.item.delete
//
// 删除单条商品
type TaobaoitemdeleteAPIResponse struct {
	model.CommonResponse
	TaobaoitemdeleteAPIResponseModel
}

// TaobaoitemdeleteAPIResponseModel is 删除单条商品 成功返回结果
type TaobaoitemdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"item_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 被删除商品的相关信息
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}
