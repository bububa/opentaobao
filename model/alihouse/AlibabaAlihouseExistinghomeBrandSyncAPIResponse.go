package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrandsyncAPIResponse 二手房公司品牌数据同步 API返回值
// alibaba.alihouse.existinghome.brand.sync
//
// 二手房公司品牌数据同步
type AlibabaalihouseexistinghomebrandsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomebrandsyncAPIResponseModel
}

// AlibabaalihouseexistinghomebrandsyncAPIResponseModel is 二手房公司品牌数据同步 成功返回结果
type AlibabaalihouseexistinghomebrandsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_brand_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomebrandsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
