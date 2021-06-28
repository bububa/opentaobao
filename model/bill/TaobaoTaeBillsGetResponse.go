package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询账单明细 APIResponse
taobao.tae.bills.get

tae查询账单明细
*/
type TaobaoTaeBillsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tae_bills_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 当前页查询返回的结果数(0-100)。相同的查询时间段条件下，最大只能获取总共5000条记录。所以当大于等于5000时 ISV可以通过start_time及end_time来进行拆分，以保证可以查询到全部数据
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"