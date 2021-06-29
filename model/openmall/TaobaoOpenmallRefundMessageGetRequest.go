package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openmall获取退款单留言 APIRequest
taobao.openmall.refund.message.get

openmall获取退款单留言
*/
type TaobaoOpenmallRefundMessageGetRequest struct {
    model.Params

    // 分销者身份
    distributor   string 

    // 翻页页码
    pageNo   int64 

    // 翻页大小
    pageSize   int64 

    // 退款单号
    refundId   int64 

}

func NewTaobaoOpenmallRefundMessageGetRequest() *TaobaoOpenmallRefundMessageGetRequest{
    return &TaobaoOpenmallRefundMessageGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallRefundMessageGetRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.message.get"
}

func (r TaobaoOpenmallRefundMessageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallRefundMessageGetRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallRefundMessageGetRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallRefundMessageGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoOpenmallRefundMessageGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoOpenmallRefundMessageGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoOpenmallRefundMessageGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoOpenmallRefundMessageGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoOpenmallRefundMessageGetRequest) GetRefundId() int64 {
    return r.refundId
}

