package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsaddresssearchAPIResponse 查询卖家地址库 API返回值
// taobao.logistics.address.search
//
// 通过此接口查询卖家地址库，
type TaobaologisticsaddresssearchAPIResponse struct {
	model.CommonResponse
	TaobaologisticsaddresssearchAPIResponseModel
}

// TaobaologisticsaddresssearchAPIResponseModel is 查询卖家地址库 成功返回结果
type TaobaologisticsaddresssearchAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_address_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 一组地址库数据，
	Addresses []AddressResult `json:"addresses,omitempty" xml:"addresses>address_result,omitempty"`
}
