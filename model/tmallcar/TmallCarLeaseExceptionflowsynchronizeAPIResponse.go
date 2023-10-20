package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleaseexceptionflowsynchronizeAPIResponse 天猫开新车租后异常流线下处理状态通知接口 API返回值
// tmall.car.lease.exceptionflowsynchronize
//
// 天猫开新车租后异常流线下处理状态通知接口
type TmallcarleaseexceptionflowsynchronizeAPIResponse struct {
	model.CommonResponse
	TmallcarleaseexceptionflowsynchronizeAPIResponseModel
}

// TmallcarleaseexceptionflowsynchronizeAPIResponseModel is 天猫开新车租后异常流线下处理状态通知接口 成功返回结果
type TmallcarleaseexceptionflowsynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_exceptionflowsynchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
