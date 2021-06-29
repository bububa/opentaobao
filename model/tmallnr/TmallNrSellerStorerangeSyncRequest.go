package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步商户中心服务范围 API请求
tmall.nr.seller.storerange.sync

同步商户中心服务范围
*/
type TmallNrSellerStorerangeSyncRequest struct {
    model.Params
    // 业务身份标识,dss定时送；self_day 自配日达；self_hour 自配小时达
    bizIdentity   string
    // 系统自动生成
    reqDTOList   []SyncServiceRangeRequestDto
    // 卖家id，有可能和登录seller不是同一个id
    sellerId   int64
}

// 初始化TmallNrSellerStorerangeSyncRequest对象
func NewTmallNrSellerStorerangeSyncRequest() *TmallNrSellerStorerangeSyncRequest{
    return &TmallNrSellerStorerangeSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrSellerStorerangeSyncRequest) GetApiMethodName() string {
    return "tmall.nr.seller.storerange.sync"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrSellerStorerangeSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizIdentity Setter
// 业务身份标识,dss定时送；self_day 自配日达；self_hour 自配小时达
func (r *TmallNrSellerStorerangeSyncRequest) SetBizIdentity(bizIdentity string) error {
    r.bizIdentity = bizIdentity
    r.Set("biz_identity", bizIdentity)
    return nil
}

// BizIdentity Getter
func (r TmallNrSellerStorerangeSyncRequest) GetBizIdentity() string {
    return r.bizIdentity
}
// ReqDTOList Setter
// 系统自动生成
func (r *TmallNrSellerStorerangeSyncRequest) SetReqDTOList(reqDTOList []SyncServiceRangeRequestDto) error {
    r.reqDTOList = reqDTOList
    r.Set("req_d_t_o_list", reqDTOList)
    return nil
}

// ReqDTOList Getter
func (r TmallNrSellerStorerangeSyncRequest) GetReqDTOList() []SyncServiceRangeRequestDto {
    return r.reqDTOList
}
// SellerId Setter
// 卖家id，有可能和登录seller不是同一个id
func (r *TmallNrSellerStorerangeSyncRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TmallNrSellerStorerangeSyncRequest) GetSellerId() int64 {
    return r.sellerId
}
