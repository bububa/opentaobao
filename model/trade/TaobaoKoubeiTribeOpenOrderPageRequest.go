package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑综合体订单列表信息查询 APIRequest
taobao.koubei.tribe.open.order.page

查询口碑商圈用户的订单列表信息
*/
type TaobaoKoubeiTribeOpenOrderPageRequest struct {
    model.Params

    // 订单状态；ALL（全部），WAIT_PAY（代付款），WAIT_CONSUME（代消费）
    orderStatus   string 

    // 每页大小
    pageSize   int64 

    // 起始页
    pageNo   int64 

    // 数据集Id
    dataSetId   string 

    // 用户openId
    openId   string 

}

func NewTaobaoKoubeiTribeOpenOrderPageRequest() *TaobaoKoubeiTribeOpenOrderPageRequest{
    return &TaobaoKoubeiTribeOpenOrderPageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetApiMethodName() string {
    return "taobao.koubei.tribe.open.order.page"
}

func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetOrderStatus(orderStatus string) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetOrderStatus() string {
    return r.orderStatus
}

func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetOpenId() string {
    return r.openId
}

