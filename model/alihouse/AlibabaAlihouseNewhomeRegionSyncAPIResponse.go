package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeregionsyncAPIResponse 城区数据同步 API返回值
// alibaba.alihouse.newhome.region.sync
//
// 城区数据同步
type AlibabaalihousenewhomeregionsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeregionsyncAPIResponseModel
}

// AlibabaalihousenewhomeregionsyncAPIResponseModel is 城区数据同步 成功返回结果
type AlibabaalihousenewhomeregionsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_region_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeregionsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
