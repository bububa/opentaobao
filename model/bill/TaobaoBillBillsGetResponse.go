package bill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询账单明细数据(自研发商家专用) APIResponse
taobao.bill.bills.get

查询账单明细数据
*/
type TaobaoBillBillsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBillBillsGetResponse `json:"taobao_bill_bills_get_response,omitempty"`
}

type TaobaoBillBillsGetResponse struct {

    // 账单列表
    Bills   []Bill `json:"bills,omitempty"`

    // 是否存在下一页
    HasNext   bool `json:"has_next,omitempty"`

    // 当前页查询返回的结果数(0-100)。相同的查询时间段条件下，最大只能获取总共5000条记录。所以当大于等于5000时 ISV可以通过start_time及end_time来进行拆分，以保证可以查询到全部数据
    TotalResults   int64 `json:"total_results,omitempty"`

}
