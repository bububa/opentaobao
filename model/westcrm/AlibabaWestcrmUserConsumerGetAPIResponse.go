package westcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawestcrmuserconsumergetAPIResponse 获取指定用户的消费总额 API返回值
// alibaba.westcrm.user.consumer.get
//
// 获取指定用户的消费总额
type AlibabawestcrmuserconsumergetAPIResponse struct {
	model.CommonResponse
	AlibabawestcrmuserconsumergetAPIResponseModel
}

// AlibabawestcrmuserconsumergetAPIResponseModel is 获取指定用户的消费总额 成功返回结果
type AlibabawestcrmuserconsumergetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_user_consumer_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
