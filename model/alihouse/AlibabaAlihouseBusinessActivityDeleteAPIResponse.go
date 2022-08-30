package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseBusinessActivityDeleteAPIResponse 电商券活动删除 API返回值
// alibaba.alihouse.business.activity.delete
//
// 电商券活动删除
type AlibabaAlihouseBusinessActivityDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseBusinessActivityDeleteAPIResponseModel
}

// AlibabaAlihouseBusinessActivityDeleteAPIResponseModel is 电商券活动删除 成功返回结果
type AlibabaAlihouseBusinessActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_business_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseBusinessActivityDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
