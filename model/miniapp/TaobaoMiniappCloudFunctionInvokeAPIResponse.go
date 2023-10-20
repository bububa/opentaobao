package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappcloudfunctioninvokeAPIResponse 外部触发云函数 API返回值
// taobao.miniapp.cloud.function.invoke
//
// 用户isv从外部触发聚石塔云函数的执行。
type TaobaominiappcloudfunctioninvokeAPIResponse struct {
	model.CommonResponse
	TaobaominiappcloudfunctioninvokeAPIResponseModel
}

// TaobaominiappcloudfunctioninvokeAPIResponseModel is 外部触发云函数 成功返回结果
type TaobaominiappcloudfunctioninvokeAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_cloud_function_invoke_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数，长度限制8个字符，超出部分会被截断
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
