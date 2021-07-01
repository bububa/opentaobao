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
type TmallNrSellerStorerangeSyncAPIRequest struct {
    model.Params
    // 业务身份标识,dss定时送；self_day 自配日达；self_hour 自配小时达
    _bizIdentity   string
    // 系统自动生成
    _reqDTOList   []SyncServiceRangeRequestDTO
    // 卖家id，有可能和登录seller不是同一个id
    _sellerId   int64
}

// 初始化TmallNrSellerStorerangeSyncAPIRequest对象
func NewTmallNrSellerStorerangeSyncRequest() *TmallNrSellerStorerangeSyncAPIRequest{
    return &TmallNrSellerStorerangeSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrSellerStorerangeSyncAPIRequest) GetApiMethodName() string {
    return "tmall.nr.seller.storerange.sync"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrSellerStorerangeSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizIdentity Setter
// 业务身份标识,dss定时送；self_day 自配日达；self_hour 自配小时达
func (r *TmallNrSellerStorerangeSyncAPIRequest) SetBizIdentity(_bizIdentity string) error {
    r._bizIdentity = _bizIdentity
    r.Set("biz_identity", _bizIdentity)
    return nil
}

// BizIdentity Getter
func (r TmallNrSellerStorerangeSyncAPIRequest) GetBizIdentity() string {
    return r._bizIdentity
}
// ReqDTOList Setter
// 系统自动生成
func (r *TmallNrSellerStorerangeSyncAPIRequest) SetReqDTOList(_reqDTOList []SyncServiceRangeRequestDTO) error {
    r._reqDTOList = _reqDTOList
    r.Set("req_d_t_o_list", _reqDTOList)
    return nil
}

// ReqDTOList Getter
func (r TmallNrSellerStorerangeSyncAPIRequest) GetReqDTOList() []SyncServiceRangeRequestDTO {
    return r._reqDTOList
}
// SellerId Setter
// 卖家id，有可能和登录seller不是同一个id
func (r *TmallNrSellerStorerangeSyncAPIRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TmallNrSellerStorerangeSyncAPIRequest) GetSellerId() int64 {
    return r._sellerId
}
