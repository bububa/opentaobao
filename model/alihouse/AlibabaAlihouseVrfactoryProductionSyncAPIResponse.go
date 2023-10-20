package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousevrfactoryproductionsyncAPIResponse vr生产数据上翻 API返回值
// alibaba.alihouse.vrfactory.production.sync
//
// vr生产数据上翻
type AlibabaalihousevrfactoryproductionsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousevrfactoryproductionsyncAPIResponseModel
}

// AlibabaalihousevrfactoryproductionsyncAPIResponseModel is vr生产数据上翻 成功返回结果
type AlibabaalihousevrfactoryproductionsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_vrfactory_production_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousevrfactoryproductionsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
