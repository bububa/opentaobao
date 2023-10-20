package uscesl

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizitemlightupAPIResponse 商品条码亮灯API API返回值
// taobao.uscesl.biz.item.light.up
//
// 亮灯API
type TaobaousceslbizitemlightupAPIResponse struct {
	model.CommonResponse
	TaobaousceslbizitemlightupAPIResponseModel
}

// TaobaousceslbizitemlightupAPIResponseModel is 商品条码亮灯API 成功返回结果
type TaobaousceslbizitemlightupAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_item_light_up_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaousceslbizitemlightupResult `json:"result,omitempty" xml:"result,omitempty"`
}
