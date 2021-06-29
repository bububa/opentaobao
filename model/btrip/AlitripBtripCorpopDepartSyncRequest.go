package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部部门同步 API请求
alitrip.btrip.corpop.depart.sync

同步外部平台部门信息至商旅内部
*/
type AlitripBtripCorpopDepartSyncRequest struct {
    model.Params
    // 同步部门请求
    rq   *BtripDepartSyncRq
}

// 初始化AlitripBtripCorpopDepartSyncRequest对象
func NewAlitripBtripCorpopDepartSyncRequest() *AlitripBtripCorpopDepartSyncRequest{
    return &AlitripBtripCorpopDepartSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopDepartSyncRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.depart.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopDepartSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 同步部门请求
func (r *AlitripBtripCorpopDepartSyncRequest) SetRq(rq *BtripDepartSyncRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCorpopDepartSyncRequest) GetRq() *BtripDepartSyncRq {
    return r.rq
}
