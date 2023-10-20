package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousedownselfAPIResponse 房源信息下架 API返回值
// alibaba.alihouse.existinghome.house.downself
//
// 房源信息下架
type AlibabaalihouseexistinghomehousedownselfAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomehousedownselfAPIResponseModel
}

// AlibabaalihouseexistinghomehousedownselfAPIResponseModel is 房源信息下架 成功返回结果
type AlibabaalihouseexistinghomehousedownselfAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_downself_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomehousedownselfResult `json:"result,omitempty" xml:"result,omitempty"`
}
