package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取默认状态下商品列表 API请求
alibaba.mos.store.getdefautitems

获取默认状态下商品列表
*/
type AlibabaMosStoreGetdefautitemsRequest struct {
    model.Params
    // 屏编号
    screenNo   string
    // 分页查询开始index
    start   int64
    // 分页查询每页记录数
    limitCount   int64
}

// 初始化AlibabaMosStoreGetdefautitemsRequest对象
func NewAlibabaMosStoreGetdefautitemsRequest() *AlibabaMosStoreGetdefautitemsRequest{
    return &AlibabaMosStoreGetdefautitemsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosStoreGetdefautitemsRequest) GetApiMethodName() string {
    return "alibaba.mos.store.getdefautitems"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosStoreGetdefautitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScreenNo Setter
// 屏编号
func (r *AlibabaMosStoreGetdefautitemsRequest) SetScreenNo(screenNo string) error {
    r.screenNo = screenNo
    r.Set("screen_no", screenNo)
    return nil
}

// ScreenNo Getter
func (r AlibabaMosStoreGetdefautitemsRequest) GetScreenNo() string {
    return r.screenNo
}
// Start Setter
// 分页查询开始index
func (r *AlibabaMosStoreGetdefautitemsRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

// Start Getter
func (r AlibabaMosStoreGetdefautitemsRequest) GetStart() int64 {
    return r.start
}
// LimitCount Setter
// 分页查询每页记录数
func (r *AlibabaMosStoreGetdefautitemsRequest) SetLimitCount(limitCount int64) error {
    r.limitCount = limitCount
    r.Set("limit_count", limitCount)
    return nil
}

// LimitCount Getter
func (r AlibabaMosStoreGetdefautitemsRequest) GetLimitCount() int64 {
    return r.limitCount
}
