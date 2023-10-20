package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobanamadpcitemupdateAPIResponse 编辑商品 API返回值
// taobao.banamadpc.item.update
//
// 巴拿马供应商通过此接口编辑商品
type TaobaobanamadpcitemupdateAPIResponse struct {
	model.CommonResponse
	TaobaobanamadpcitemupdateAPIResponseModel
}

// TaobaobanamadpcitemupdateAPIResponseModel is 编辑商品 成功返回结果
type TaobaobanamadpcitemupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"banamadpc_item_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	ApiResult *TaobaobanamadpcitemupdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
