package jstinteractive

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractivePointIncreaseAPIResponse 互动积分发放接口 API返回值
// taobao.jst.interactive.point.increase
//
// 向用户发放互动积分
type TaobaoJstInteractivePointIncreaseAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractivePointIncreaseAPIResponseModel
}

// TaobaoJstInteractivePointIncreaseAPIResponseModel is 互动积分发放接口 成功返回结果
type TaobaoJstInteractivePointIncreaseAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_point_increase_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户积分总额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
