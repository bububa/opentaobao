package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIResponse 门店委托信息查询 API返回值
// alibaba.alihouse.existinghome.house.storeentrust.query
//
// 门店委托信息查询
type AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIResponseModel
}

// AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIResponseModel is 门店委托信息查询 成功返回结果
type AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_storeentrust_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseStoreentrustQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
