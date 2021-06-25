package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取默认状态下商品列表 APIRequest
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

func NewAlibabaMosStoreGetdefautitemsRequest() *AlibabaMosStoreGetdefautitemsRequest{
    return &AlibabaMosStoreGetdefautitemsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosStoreGetdefautitemsRequest) GetApiMethodName() string {
    return "alibaba.mos.store.getdefautitems"
}

func (r AlibabaMosStoreGetdefautitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosStoreGetdefautitemsRequest) SetScreenNo(screenNo string) error {
    r.screenNo = screenNo
    r.Set("screen_no", screenNo)
    return nil
}

func (r AlibabaMosStoreGetdefautitemsRequest) GetScreenNo() string {
    return r.screenNo
}

func (r *AlibabaMosStoreGetdefautitemsRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r AlibabaMosStoreGetdefautitemsRequest) GetStart() int64 {
    return r.start
}

func (r *AlibabaMosStoreGetdefautitemsRequest) SetLimitCount(limitCount int64) error {
    r.limitCount = limitCount
    r.Set("limit_count", limitCount)
    return nil
}

func (r AlibabaMosStoreGetdefautitemsRequest) GetLimitCount() int64 {
    return r.limitCount
}

