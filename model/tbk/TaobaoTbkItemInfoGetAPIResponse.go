package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkItemInfoGetAPIResponse 淘宝客-公用-淘宝客商品详情查询(简版) API返回值
// taobao.tbk.item.info.get
//
// 淘宝客商品详情查询（简版）
type TaobaoTbkItemInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkItemInfoGetAPIResponseModel
}

// TaobaoTbkItemInfoGetAPIResponseModel is 淘宝客-公用-淘宝客商品详情查询(简版) 成功返回结果
type TaobaoTbkItemInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_item_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘宝客商品
	Results []NTbkItem `json:"results,omitempty" xml:"results>n_tbk_item,omitempty"`
}
