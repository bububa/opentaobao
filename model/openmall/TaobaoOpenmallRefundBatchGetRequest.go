package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取openmall退款单 APIRequest
taobao.openmall.refund.batch.get

批量获取openmall退款单
注意：该接口信息存在延迟，如需实时详情请访问taobao.openmall.refund.get
*/
type TaobaoOpenmallRefundBatchGetRequest struct {
    model.Params

    // 查询范围结束时间，闭区间
    endCreated   string 

    // 翻页页码，从1开始
    pageIndex   int64 

    // 页面大小，不超过100
    pageSize   int64 

    // 查询的渠道商Nick
    distributor   string 

    // 查询范围开始时间，闭区间
    startCreated   string 

}

func NewTaobaoOpenmallRefundBatchGetRequest() *TaobaoOpenmallRefundBatchGetRequest{
    return &TaobaoOpenmallRefundBatchGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallRefundBatchGetRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.batch.get"
}

func (r TaobaoOpenmallRefundBatchGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallRefundBatchGetRequest) SetEndCreated(endCreated string) error {
    r.endCreated = endCreated
    r.Set("end_created", endCreated)
    return nil
}

func (r TaobaoOpenmallRefundBatchGetRequest) GetEndCreated() string {
    return r.endCreated
}

func (r *TaobaoOpenmallRefundBatchGetRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoOpenmallRefundBatchGetRequest) GetPageIndex() int64 {
    return r.pageIndex
}

func (r *TaobaoOpenmallRefundBatchGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoOpenmallRefundBatchGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoOpenmallRefundBatchGetRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallRefundBatchGetRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallRefundBatchGetRequest) SetStartCreated(startCreated string) error {
    r.startCreated = startCreated
    r.Set("start_created", startCreated)
    return nil
}

func (r TaobaoOpenmallRefundBatchGetRequest) GetStartCreated() string {
    return r.startCreated
}

