package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse 委托房源变更标准房源 API返回值
// alibaba.alihouse.existinghome.house.change.standard
//
// 委托房源变更标准房源
type AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponseModel
}

// AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponseModel is 委托房源变更标准房源 成功返回结果
type AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_change_standard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseChangeStandardResult `json:"result,omitempty" xml:"result,omitempty"`
}
