package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrandcitysyncAPIResponse 二手房城市品牌店数据同步 API返回值
// alibaba.alihouse.existinghome.brandcity.sync
//
// 二手房城市品牌店数据同步
type AlibabaalihouseexistinghomebrandcitysyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomebrandcitysyncAPIResponseModel
}

// AlibabaalihouseexistinghomebrandcitysyncAPIResponseModel is 二手房城市品牌店数据同步 成功返回结果
type AlibabaalihouseexistinghomebrandcitysyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_brandcity_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomebrandcitysyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
