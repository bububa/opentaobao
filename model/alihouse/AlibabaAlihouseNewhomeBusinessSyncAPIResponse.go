package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomebusinesssyncAPIResponse 商圈数据同步 API返回值
// alibaba.alihouse.newhome.business.sync
//
// 商圈数据同步
type AlibabaalihousenewhomebusinesssyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomebusinesssyncAPIResponseModel
}

// AlibabaalihousenewhomebusinesssyncAPIResponseModel is 商圈数据同步 成功返回结果
type AlibabaalihousenewhomebusinesssyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_business_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomebusinesssyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
