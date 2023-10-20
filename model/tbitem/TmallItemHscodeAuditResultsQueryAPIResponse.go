package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemHscodeAuditResultsQueryAPIResponse 商品hscode信息审核状态查询接口 API返回值
// tmall.item.hscode.audit.results.query
//
// 通过此接口查询天猫跨境商品的hscode信息审核状态，卖家可以参考返回结果判断是否需要调整商品hscode相关信息。
type TmallItemHscodeAuditResultsQueryAPIResponse struct {
	model.CommonResponse
	TmallItemHscodeAuditResultsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemHscodeAuditResultsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemHscodeAuditResultsQueryAPIResponseModel).Reset()
}

// TmallItemHscodeAuditResultsQueryAPIResponseModel is 商品hscode信息审核状态查询接口 成功返回结果
type TmallItemHscodeAuditResultsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_hscode_audit_results_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品或sku的hscode信息审核状态。
	ResultList []HscodeAuditInfo `json:"result_list,omitempty" xml:"result_list>hscode_audit_info,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemHscodeAuditResultsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolTmallItemHscodeAuditResultsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemHscodeAuditResultsQueryAPIResponse)
	},
}

// GetTmallItemHscodeAuditResultsQueryAPIResponse 从 sync.Pool 获取 TmallItemHscodeAuditResultsQueryAPIResponse
func GetTmallItemHscodeAuditResultsQueryAPIResponse() *TmallItemHscodeAuditResultsQueryAPIResponse {
	return poolTmallItemHscodeAuditResultsQueryAPIResponse.Get().(*TmallItemHscodeAuditResultsQueryAPIResponse)
}

// ReleaseTmallItemHscodeAuditResultsQueryAPIResponse 将 TmallItemHscodeAuditResultsQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallItemHscodeAuditResultsQueryAPIResponse(v *TmallItemHscodeAuditResultsQueryAPIResponse) {
	v.Reset()
	poolTmallItemHscodeAuditResultsQueryAPIResponse.Put(v)
}
