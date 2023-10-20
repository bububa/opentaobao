package cainiaohandover

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaogloballogisticscarrierquerylistAPIResponse 实际承运商查询 API返回值
// cainiao.global.logistics.carrier.querylist
//
// 查询出所有的实际承运商
type CainiaogloballogisticscarrierquerylistAPIResponse struct {
	model.CommonResponse
	CainiaogloballogisticscarrierquerylistAPIResponseModel
}

// CainiaogloballogisticscarrierquerylistAPIResponseModel is 实际承运商查询 成功返回结果
type CainiaogloballogisticscarrierquerylistAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_logistics_carrier_querylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *DubboResult `json:"result,omitempty" xml:"result,omitempty"`
}
