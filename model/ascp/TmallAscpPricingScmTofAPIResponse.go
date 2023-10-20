package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAscpPricingScmTofAPIResponse TOF&SCM营销域对接-成本录入设置 API返回值
// tmall.ascp.pricing.scm.tof
//
// TOF&amp;SCM营销域对接-成本录入设置
type TmallAscpPricingScmTofAPIResponse struct {
	model.CommonResponse
	TmallAscpPricingScmTofAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAscpPricingScmTofAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAscpPricingScmTofAPIResponseModel).Reset()
}

// TmallAscpPricingScmTofAPIResponseModel is TOF&SCM营销域对接-成本录入设置 成功返回结果
type TmallAscpPricingScmTofAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_ascp_pricing_scm_tof_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// reasonsForPartSucc
	ReasonsForPartSuccList []string `json:"reasons_for_part_succ_list,omitempty" xml:"reasons_for_part_succ_list>string,omitempty"`
	// 成功
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 成功
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *TmallAscpPricingScmTofAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReasonsForPartSuccList = m.ReasonsForPartSuccList[:0]
	m.ResultMsg = ""
	m.ResultCode = ""
}

var poolTmallAscpPricingScmTofAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAscpPricingScmTofAPIResponse)
	},
}

// GetTmallAscpPricingScmTofAPIResponse 从 sync.Pool 获取 TmallAscpPricingScmTofAPIResponse
func GetTmallAscpPricingScmTofAPIResponse() *TmallAscpPricingScmTofAPIResponse {
	return poolTmallAscpPricingScmTofAPIResponse.Get().(*TmallAscpPricingScmTofAPIResponse)
}

// ReleaseTmallAscpPricingScmTofAPIResponse 将 TmallAscpPricingScmTofAPIResponse 保存到 sync.Pool
func ReleaseTmallAscpPricingScmTofAPIResponse(v *TmallAscpPricingScmTofAPIResponse) {
	v.Reset()
	poolTmallAscpPricingScmTofAPIResponse.Put(v)
}
