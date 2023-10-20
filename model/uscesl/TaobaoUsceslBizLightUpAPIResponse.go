package uscesl

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizlightupAPIResponse 价签LED等点亮 API返回值
// taobao.uscesl.biz.light.up
//
// 价签LED等点亮
type TaobaousceslbizlightupAPIResponse struct {
	model.CommonResponse
	TaobaousceslbizlightupAPIResponseModel
}

// TaobaousceslbizlightupAPIResponseModel is 价签LED等点亮 成功返回结果
type TaobaousceslbizlightupAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_light_up_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaousceslbizlightupResult `json:"result,omitempty" xml:"result,omitempty"`
}
