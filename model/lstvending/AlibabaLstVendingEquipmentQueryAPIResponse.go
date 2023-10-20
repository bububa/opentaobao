package lstvending

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendingequipmentqueryAPIResponse 自动售卖机设备信息查询 API返回值
// alibaba.lst.vending.equipment.query
//
// 为零售通品牌商提供已租赁的设备信息查询。
type AlibabalstvendingequipmentqueryAPIResponse struct {
	model.CommonResponse
	AlibabalstvendingequipmentqueryAPIResponseModel
}

// AlibabalstvendingequipmentqueryAPIResponseModel is 自动售卖机设备信息查询 成功返回结果
type AlibabalstvendingequipmentqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vending_equipment_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *AlibabalstvendingequipmentqueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
