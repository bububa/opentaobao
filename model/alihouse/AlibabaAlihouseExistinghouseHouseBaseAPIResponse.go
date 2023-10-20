package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghousehousebaseAPIResponse （租房）同步房屋信息 API返回值
// alibaba.alihouse.existinghouse.house.base
//
// 房屋信息上翻
type AlibabaalihouseexistinghousehousebaseAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghousehousebaseAPIResponseModel
}

// AlibabaalihouseexistinghousehousebaseAPIResponseModel is （租房）同步房屋信息 成功返回结果
type AlibabaalihouseexistinghousehousebaseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghouse_house_base_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaalihouseexistinghousehousebaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
