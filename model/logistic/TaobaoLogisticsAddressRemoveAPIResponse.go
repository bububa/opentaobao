package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsaddressremoveAPIResponse 删除卖家地址库 API返回值
// taobao.logistics.address.remove
//
// 用此接口删除卖家地址库
type TaobaologisticsaddressremoveAPIResponse struct {
	model.CommonResponse
	TaobaologisticsaddressremoveAPIResponseModel
}

// TaobaologisticsaddressremoveAPIResponseModel is 删除卖家地址库 成功返回结果
type TaobaologisticsaddressremoveAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_address_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 只返回修改日期modify_date
	AddressResult *AddressResult `json:"address_result,omitempty" xml:"address_result,omitempty"`
}
