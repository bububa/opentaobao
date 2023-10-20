package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseUpselfAPIResponse 房源上架 API返回值
// alibaba.alihouse.existinghome.house.upself
//
// 房源信息上架
type AlibabaAlihouseExistinghomeHouseUpselfAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseUpselfAPIResponseModel
}

// AlibabaAlihouseExistinghomeHouseUpselfAPIResponseModel is 房源上架 成功返回结果
type AlibabaAlihouseExistinghomeHouseUpselfAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_upself_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseUpselfResult `json:"result,omitempty" xml:"result,omitempty"`
}
