package bill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeBillsGetAPIResponse tae查询账单明细 API返回值
// taobao.tae.bills.get
//
// tae查询账单明细
type TaobaoTaeBillsGetAPIResponse struct {
	model.CommonResponse
	TaobaoTaeBillsGetAPIResponseModel
}

// TaobaoTaeBillsGetAPIResponseModel is tae查询账单明细 成功返回结果
type TaobaoTaeBillsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tae_bills_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 当前页查询返回的结果数(0-100)。相同的查询时间段条件下，最大只能获取总共5000条记录。所以当大于等于5000时 ISV可以通过start_time及end_time来进行拆分，以保证可以查询到全部数据
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 账单列表
	Bills []BillDto `json:"bills,omitempty" xml:"bills>bill_dto,omitempty"`
}
