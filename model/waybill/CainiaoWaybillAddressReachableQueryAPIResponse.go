package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillAddressReachableQueryAPIResponse 地址可达查询 API返回值
// cainiao.waybill.address.reachable.query
//
// 地址可达查询
type CainiaoWaybillAddressReachableQueryAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillAddressReachableQueryAPIResponseModel
}

// CainiaoWaybillAddressReachableQueryAPIResponseModel is 地址可达查询 成功返回结果
type CainiaoWaybillAddressReachableQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_address_reachable_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
