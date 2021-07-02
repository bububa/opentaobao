package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAddressReachableAPIResponse 判定服务是否可达 API返回值
// taobao.logistics.address.reachable
//
// 根据输入的目标地址，判断服务是否可达。
// 现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
type TaobaoLogisticsAddressReachableAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsAddressReachableAPIResponseModel
}

// TaobaoLogisticsAddressReachableAPIResponseModel is 判定服务是否可达 成功返回结果
type TaobaoLogisticsAddressReachableAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_address_reachable_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 地址可达返回结果，每个TP对应一个
	ReachableResultList []AddressReachableResult `json:"reachable_result_list,omitempty" xml:"reachable_result_list>address_reachable_result,omitempty"`
}
