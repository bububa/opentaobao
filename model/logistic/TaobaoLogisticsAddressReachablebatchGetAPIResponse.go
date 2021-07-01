package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsAddressReachablebatchGetAPIResponse
批量判定服务是否可达 API返回值
taobao.logistics.address.reachablebatch.get

批量判定服务是否可达 */
type TaobaoLogisticsAddressReachablebatchGetAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsAddressReachablebatchGetAPIResponseModel
}

// TaobaoLogisticsAddressReachablebatchGetAPIResponseModel is 批量判定服务是否可达 成功返回结果
type TaobaoLogisticsAddressReachablebatchGetAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_address_reachablebatch_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流是否可达结果列表
	ReachableResults []AddressReachableTopResult `json:"reachable_results,omitempty" xml:"reachable_results>address_reachable_top_result,omitempty"`
}
