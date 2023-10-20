package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptCustbaseGetAPIResponse 客户账户报表基础数据对象 API返回值
// taobao.simba.rpt.custbase.get
//
// 客户账户报表基础数据对象
type TaobaoSimbaRptCustbaseGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptCustbaseGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptCustbaseGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptCustbaseGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptCustbaseGetAPIResponseModel is 客户账户报表基础数据对象 成功返回结果
type TaobaoSimbaRptCustbaseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_custbase_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户帐户结果
	RptCustBaseList string `json:"rpt_cust_base_list,omitempty" xml:"rpt_cust_base_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptCustbaseGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RptCustBaseList = ""
}

var poolTaobaoSimbaRptCustbaseGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptCustbaseGetAPIResponse)
	},
}

// GetTaobaoSimbaRptCustbaseGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptCustbaseGetAPIResponse
func GetTaobaoSimbaRptCustbaseGetAPIResponse() *TaobaoSimbaRptCustbaseGetAPIResponse {
	return poolTaobaoSimbaRptCustbaseGetAPIResponse.Get().(*TaobaoSimbaRptCustbaseGetAPIResponse)
}

// ReleaseTaobaoSimbaRptCustbaseGetAPIResponse 将 TaobaoSimbaRptCustbaseGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptCustbaseGetAPIResponse(v *TaobaoSimbaRptCustbaseGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptCustbaseGetAPIResponse.Put(v)
}
