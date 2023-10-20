package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomemetrosyncAPIResponse 地铁数据同步 API返回值
// alibaba.alihouse.newhome.metro.sync
//
// 地铁数据同步
type AlibabaalihousenewhomemetrosyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomemetrosyncAPIResponseModel
}

// AlibabaalihousenewhomemetrosyncAPIResponseModel is 地铁数据同步 成功返回结果
type AlibabaalihousenewhomemetrosyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_metro_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomemetrosyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
