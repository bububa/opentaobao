package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivitySaveAPIResponse 新增或者更新销售活动 API返回值
// alibaba.alihouse.newhome.activity.save
//
// 新增或者更新销售活动
type AlibabaAlihouseNewhomeActivitySaveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeActivitySaveAPIResponseModel
}

// AlibabaAlihouseNewhomeActivitySaveAPIResponseModel is 新增或者更新销售活动 成功返回结果
type AlibabaAlihouseNewhomeActivitySaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_activity_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeActivitySaveResult `json:"result,omitempty" xml:"result,omitempty"`
}
