package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取openmall订单 APIRequest
taobao.openmall.trade.batch.get

批量获取openmall订单
注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取
*/
type TaobaoOpenmallTradeBatchGetRequest struct {
    model.Params

    // 查询范围结束时间，闭区间
    endCreated   string 

    // 查询页码，从1开始
    pageIndex   int64 

    // 页面大小，不超过100
    pageSize   int64 

    // 渠道商Nick
    distributor   string 

    // 查询范围开始时间，闭区间
    startCreated   string 

}

func NewTaobaoOpenmallTradeBatchGetRequest() *TaobaoOpenmallTradeBatchGetRequest{
    return &TaobaoOpenmallTradeBatchGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallTradeBatchGetRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.batch.get"
}

func (r TaobaoOpenmallTradeBatchGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallTradeBatchGetRequest) SetEndCreated(endCreated string) error {
    r.endCreated = endCreated
    r.Set("end_created", endCreated)
    return nil
}

func (r TaobaoOpenmallTradeBatchGetRequest) GetEndCreated() string {
    return r.endCreated
}

func (r *TaobaoOpenmallTradeBatchGetRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoOpenmallTradeBatchGetRequest) GetPageIndex() int64 {
    return r.pageIndex
}

func (r *TaobaoOpenmallTradeBatchGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoOpenmallTradeBatchGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoOpenmallTradeBatchGetRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallTradeBatchGetRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallTradeBatchGetRequest) SetStartCreated(startCreated string) error {
    r.startCreated = startCreated
    r.Set("start_created", startCreated)
    return nil
}

func (r TaobaoOpenmallTradeBatchGetRequest) GetStartCreated() string {
    return r.startCreated
}

