package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomecompanysyncAPIResponse 二手房公司同步接口 API返回值
// alibaba.alihouse.existinghome.company.sync
//
// 二手房公司同步接口
type AlibabaalihouseexistinghomecompanysyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomecompanysyncAPIResponseModel
}

// AlibabaalihouseexistinghomecompanysyncAPIResponseModel is 二手房公司同步接口 成功返回结果
type AlibabaalihouseexistinghomecompanysyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_company_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomecompanysyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
