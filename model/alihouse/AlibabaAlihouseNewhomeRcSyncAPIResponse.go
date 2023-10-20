package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomercsyncAPIResponse 阿里房产图文草稿信息同步 API返回值
// alibaba.alihouse.newhome.rc.sync
//
// 接收图文草稿信息
type AlibabaalihousenewhomercsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomercsyncAPIResponseModel
}

// AlibabaalihousenewhomercsyncAPIResponseModel is 阿里房产图文草稿信息同步 成功返回结果
type AlibabaalihousenewhomercsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_rc_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomercsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
