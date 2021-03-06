package traderate

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFliggyWrateGetmixratelistAPIResponse 飞猪通用评价接口 API返回值
// taobao.fliggy.wrate.getmixratelist
//
// 飞猪评价通用接口
type TaobaoFliggyWrateGetmixratelistAPIResponse struct {
	model.CommonResponse
	TaobaoFliggyWrateGetmixratelistAPIResponseModel
}

// TaobaoFliggyWrateGetmixratelistAPIResponseModel is 飞猪通用评价接口 成功返回结果
type TaobaoFliggyWrateGetmixratelistAPIResponseModel struct {
	XMLName xml.Name `xml:"fliggy_wrate_getmixratelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFliggyWrateGetmixratelistResult `json:"result,omitempty" xml:"result,omitempty"`
}
