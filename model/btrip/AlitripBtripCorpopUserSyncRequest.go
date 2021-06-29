package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部人员同步 API请求
alitrip.btrip.corpop.user.sync

同步外部平台用户信息至商旅内部
*/
type AlitripBtripCorpopUserSyncRequest struct {
    model.Params
    // 人员同步请求
    rq   *BtripUserSyncRq
}

// 初始化AlitripBtripCorpopUserSyncRequest对象
func NewAlitripBtripCorpopUserSyncRequest() *AlitripBtripCorpopUserSyncRequest{
    return &AlitripBtripCorpopUserSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopUserSyncRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.user.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopUserSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 人员同步请求
func (r *AlitripBtripCorpopUserSyncRequest) SetRq(rq *BtripUserSyncRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCorpopUserSyncRequest) GetRq() *BtripUserSyncRq {
    return r.rq
}
