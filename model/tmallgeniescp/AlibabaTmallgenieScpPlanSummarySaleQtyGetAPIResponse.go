package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplansummarysaleqtygetAPIResponse 同步销售数据按照渠道类型汇总 API返回值
// alibaba.tmallgenie.scp.plan.summary.sale.qty.get
//
// 同步销售数据按照渠道类型汇总
type AlibabatmallgeniescpplansummarysaleqtygetAPIResponse struct {
	model.CommonResponse
	AlibabatmallgeniescpplansummarysaleqtygetAPIResponseModel
}

// AlibabatmallgeniescpplansummarysaleqtygetAPIResponseModel is 同步销售数据按照渠道类型汇总 成功返回结果
type AlibabatmallgeniescpplansummarysaleqtygetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_summary_sale_qty_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
