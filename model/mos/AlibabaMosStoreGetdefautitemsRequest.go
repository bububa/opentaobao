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
    _screenNo   string
    // 分页查询开始index
    _start   int64
    // 分页查询每页记录数
    _limitCount   int64
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
func (r *AlibabaMosStoreGetdefautitemsRequest) SetScreenNo(_screenNo string) error {
    r._screenNo = _screenNo
    r.Set("screen_no", _screenNo)
    return nil
}

// ScreenNo Getter
func (r AlibabaMosStoreGetdefautitemsRequest) GetScreenNo() string {
    return r._screenNo
}
// Start Setter
// 分页查询开始index
func (r *AlibabaMosStoreGetdefautitemsRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r AlibabaMosStoreGetdefautitemsRequest) GetStart() int64 {
    return r._start
}
// LimitCount Setter
// 分页查询每页记录数
func (r *AlibabaMosStoreGetdefautitemsRequest) SetLimitCount(_limitCount int64) error {
    r._limitCount = _limitCount
    r.Set("limit_count", _limitCount)
    return nil
}

// LimitCount Getter
func (r AlibabaMosStoreGetdefautitemsRequest) GetLimitCount() int64 {
    return r._limitCount
}
