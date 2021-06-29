package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店服务范围读取 APIRequest
tmall.nr.seller.storerange.read

读取卖家所属门店的服务范围
*/
type TmallNrSellerStorerangeReadRequest struct {
    model.Params

    // 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
    sellerId   int64 

    // 业务身份，此api非必填
    bizIdentity   string 

    // 门店id
    storeIds   []int64 

}

func NewTmallNrSellerStorerangeReadRequest() *TmallNrSellerStorerangeReadRequest{
    return &TmallNrSellerStorerangeReadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrSellerStorerangeReadRequest) GetApiMethodName() string {
    return "tmall.nr.seller.storerange.read"
}

func (r TmallNrSellerStorerangeReadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrSellerStorerangeReadRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TmallNrSellerStorerangeReadRequest) GetSellerId() int64 {
    return r.sellerId
}

func (r *TmallNrSellerStorerangeReadRequest) SetBizIdentity(bizIdentity string) error {
    r.bizIdentity = bizIdentity
    r.Set("biz_identity", bizIdentity)
    return nil
}

func (r TmallNrSellerStorerangeReadRequest) GetBizIdentity() string {
    return r.bizIdentity
}

func (r *TmallNrSellerStorerangeReadRequest) SetStoreIds(storeIds []int64) error {
    r.storeIds = storeIds
    r.Set("store_ids", storeIds)
    return nil
}

func (r TmallNrSellerStorerangeReadRequest) GetStoreIds() []int64 {
    return r.storeIds
}

