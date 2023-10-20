package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobanamadpcitemrenderAPIResponse 新发商品发布页 API返回值
// taobao.banamadpc.item.render
//
// 巴拿马供应商通过此接口新发商品发布页
type TaobaobanamadpcitemrenderAPIResponse struct {
	model.CommonResponse
	TaobaobanamadpcitemrenderAPIResponseModel
}

// TaobaobanamadpcitemrenderAPIResponseModel is 新发商品发布页 成功返回结果
type TaobaobanamadpcitemrenderAPIResponseModel struct {
	XMLName xml.Name `xml:"banamadpc_item_render_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	ApiResult *TaobaobanamadpcitemrenderApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
