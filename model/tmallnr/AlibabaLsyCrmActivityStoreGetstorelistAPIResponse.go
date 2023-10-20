package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivitystoregetstorelistAPIResponse ISV查询门店 API返回值
// alibaba.lsy.crm.activity.store.getstorelist
//
// ISV查询门店
type AlibabalsycrmactivitystoregetstorelistAPIResponse struct {
	model.CommonResponse
	AlibabalsycrmactivitystoregetstorelistAPIResponseModel
}

// AlibabalsycrmactivitystoregetstorelistAPIResponseModel is ISV查询门店 成功返回结果
type AlibabalsycrmactivitystoregetstorelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_store_getstorelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果集
	PageResultDO *PageResultDo `json:"page_result_d_o,omitempty" xml:"page_result_d_o,omitempty"`
}
