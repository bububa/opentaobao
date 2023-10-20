package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityDataUpdateAPIResponse 私域导购数据回流接口 API返回值
// alibaba.lsy.crm.activity.data.update
//
// 私域导购数据回流接口
type AlibabaLsyCrmActivityDataUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmActivityDataUpdateAPIResponseModel
}

// AlibabaLsyCrmActivityDataUpdateAPIResponseModel is 私域导购数据回流接口 成功返回结果
type AlibabaLsyCrmActivityDataUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_data_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaLsyCrmActivityDataUpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
