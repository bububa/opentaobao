package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店服务范围读取 API请求
tmall.nr.seller.storerange.read

读取卖家所属门店的服务范围
*/
type TmallNrSellerStorerangeReadRequest struct {
    model.Params
    // 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
    _sellerId   int64
    // 业务身份，此api非必填
    _bizIdentity   string
    // 门店id
    _storeIds   []int64
}

// 初始化TmallNrSellerStorerangeReadRequest对象
func NewTmallNrSellerStorerangeReadRequest() *TmallNrSellerStorerangeReadRequest{
    return &TmallNrSellerStorerangeReadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrSellerStorerangeReadRequest) GetApiMethodName() string {
    return "tmall.nr.seller.storerange.read"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrSellerStorerangeReadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerId Setter
// 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
func (r *TmallNrSellerStorerangeReadRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TmallNrSellerStorerangeReadRequest) GetSellerId() int64 {
    return r._sellerId
}
// BizIdentity Setter
// 业务身份，此api非必填
func (r *TmallNrSellerStorerangeReadRequest) SetBizIdentity(_bizIdentity string) error {
    r._bizIdentity = _bizIdentity
    r.Set("biz_identity", _bizIdentity)
    return nil
}

// BizIdentity Getter
func (r TmallNrSellerStorerangeReadRequest) GetBizIdentity() string {
    return r._bizIdentity
}
// StoreIds Setter
// 门店id
func (r *TmallNrSellerStorerangeReadRequest) SetStoreIds(_storeIds []int64) error {
    r._storeIds = _storeIds
    r.Set("store_ids", _storeIds)
    return nil
}

// StoreIds Getter
func (r TmallNrSellerStorerangeReadRequest) GetStoreIds() []int64 {
    return r._storeIds
}