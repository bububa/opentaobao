package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeentrustsellingupdateAPIResponse 管家状态及房源信息接口 API返回值
// alibaba.alihouse.existinghome.entrustselling.update
//
// 管家状态及房源信息接口
type AlibabaalihouseexistinghomeentrustsellingupdateAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomeentrustsellingupdateAPIResponseModel
}

// AlibabaalihouseexistinghomeentrustsellingupdateAPIResponseModel is 管家状态及房源信息接口 成功返回结果
type AlibabaalihouseexistinghomeentrustsellingupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_entrustselling_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomeentrustsellingupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
