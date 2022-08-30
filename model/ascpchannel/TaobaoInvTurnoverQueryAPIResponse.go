package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInvTurnoverQueryAPIResponse 业务库存流水查询 API返回值
// taobao.inv.turnover.query
//
// 业务库存流水
type TaobaoInvTurnoverQueryAPIResponse struct {
	model.CommonResponse
	TaobaoInvTurnoverQueryAPIResponseModel
}

// TaobaoInvTurnoverQueryAPIResponseModel is 业务库存流水查询 成功返回结果
type TaobaoInvTurnoverQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"inv_turnover_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoInvTurnoverQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
