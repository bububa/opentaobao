package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIResponse 查询房源状态接口 API返回值
// alibaba.alihouse.existinghome.house.trade.query.status
//
// 查询房源状态接口
type AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIResponseModel
}

// AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIResponseModel is 查询房源状态接口 成功返回结果
type AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_trade_query_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回接口对象
	Result *AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult `json:"result,omitempty" xml:"result,omitempty"`
}
