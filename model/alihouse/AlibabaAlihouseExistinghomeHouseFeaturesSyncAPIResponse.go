package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousefeaturessyncAPIResponse 房源标准打标数据同步 API返回值
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
type AlibabaalihouseexistinghomehousefeaturessyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomehousefeaturessyncAPIResponseModel
}

// AlibabaalihouseexistinghomehousefeaturessyncAPIResponseModel is 房源标准打标数据同步 成功返回结果
type AlibabaalihouseexistinghomehousefeaturessyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_features_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomehousefeaturessyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
