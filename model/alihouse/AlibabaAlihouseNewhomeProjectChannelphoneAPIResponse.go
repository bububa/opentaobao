package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectchannelphoneAPIResponse 新房渠道电话数据同步 API返回值
// alibaba.alihouse.newhome.project.channelphone
//
// 新房渠道电话数据同步
type AlibabaalihousenewhomeprojectchannelphoneAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectchannelphoneAPIResponseModel
}

// AlibabaalihousenewhomeprojectchannelphoneAPIResponseModel is 新房渠道电话数据同步 成功返回结果
type AlibabaalihousenewhomeprojectchannelphoneAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_channelphone_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectchannelphoneResult `json:"result,omitempty" xml:"result,omitempty"`
}
