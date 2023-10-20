package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousebusinessactivitydeleteAPIResponse 电商券活动删除 API返回值
// alibaba.alihouse.business.activity.delete
//
// 电商券活动删除
type AlibabaalihousebusinessactivitydeleteAPIResponse struct {
	model.CommonResponse
	AlibabaalihousebusinessactivitydeleteAPIResponseModel
}

// AlibabaalihousebusinessactivitydeleteAPIResponseModel is 电商券活动删除 成功返回结果
type AlibabaalihousebusinessactivitydeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_business_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousebusinessactivitydeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
