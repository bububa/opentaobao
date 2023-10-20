package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousedeletehouseAPIResponse 删除房源 API返回值
// alibaba.alihouse.existinghome.house.delete.house
//
// 删除房源
type AlibabaalihouseexistinghomehousedeletehouseAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomehousedeletehouseAPIResponseModel
}

// AlibabaalihouseexistinghomehousedeletehouseAPIResponseModel is 删除房源 成功返回结果
type AlibabaalihouseexistinghomehousedeletehouseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_delete_house_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomehousedeletehouseResult `json:"result,omitempty" xml:"result,omitempty"`
}
