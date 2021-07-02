package uscesl

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizApActivateAPIResponse 激活AP价签通讯模块 API返回值
// taobao.uscesl.biz.ap.activate
//
// 激活AP价签通讯模块
type TaobaoUsceslBizApActivateAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizApActivateAPIResponseModel
}

// TaobaoUsceslBizApActivateAPIResponseModel is 激活AP价签通讯模块 成功返回结果
type TaobaoUsceslBizApActivateAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_ap_activate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功与否看result.success，返回true或者false
	Result *TaobaoUsceslBizApActivateResult `json:"result,omitempty" xml:"result,omitempty"`
}
