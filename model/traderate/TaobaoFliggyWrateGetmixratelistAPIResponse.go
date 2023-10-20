package traderate

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofliggywrategetmixratelistAPIResponse 飞猪通用评价接口 API返回值
// taobao.fliggy.wrate.getmixratelist
//
// 飞猪评价通用接口
type TaobaofliggywrategetmixratelistAPIResponse struct {
	model.CommonResponse
	TaobaofliggywrategetmixratelistAPIResponseModel
}

// TaobaofliggywrategetmixratelistAPIResponseModel is 飞猪通用评价接口 成功返回结果
type TaobaofliggywrategetmixratelistAPIResponseModel struct {
	XMLName xml.Name `xml:"fliggy_wrate_getmixratelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaofliggywrategetmixratelistResult `json:"result,omitempty" xml:"result,omitempty"`
}
