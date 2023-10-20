package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresspickcodecheckAPIResponse 快递公司取货码校验 API返回值
// taobao.logistics.express.pickcode.check
//
// 快递公司取货码校验
type TaobaologisticsexpresspickcodecheckAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpresspickcodecheckAPIResponseModel
}

// TaobaologisticsexpresspickcodecheckAPIResponseModel is 快递公司取货码校验 成功返回结果
type TaobaologisticsexpresspickcodecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_pickcode_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码描述
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误码标识
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 校验成功或者异常
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
