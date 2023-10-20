package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmcustomeraddnewAPIResponse 导购域留资接口 API返回值
// alibaba.lsy.crm.customer.add.new
//
// 导购域提供留资入口
type AlibabalsycrmcustomeraddnewAPIResponse struct {
	model.CommonResponse
	AlibabalsycrmcustomeraddnewAPIResponseModel
}

// AlibabalsycrmcustomeraddnewAPIResponseModel is 导购域留资接口 成功返回结果
type AlibabalsycrmcustomeraddnewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_customer_add_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabalsycrmcustomeraddnewResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
