package kbalgo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百度批量获取本地poi接口 API请求
alibaba.kbalgo.alscpois.get

接口用于百度方获取本地生活poi数据，分页获取。
*/
type AlibabaKbalgoAlscpoisGetRequest struct {
    model.Params
    // 第几页
    _pageNum   int64
    // 每页的数量。
    _pageSize   int64
}

// 初始化AlibabaKbalgoAlscpoisGetRequest对象
func NewAlibabaKbalgoAlscpoisGetRequest() *AlibabaKbalgoAlscpoisGetRequest{
    return &AlibabaKbalgoAlscpoisGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaKbalgoAlscpoisGetRequest) GetApiMethodName() string {
    return "alibaba.kbalgo.alscpois.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaKbalgoAlscpoisGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNum Setter
// 第几页
func (r *AlibabaKbalgoAlscpoisGetRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaKbalgoAlscpoisGetRequest) GetPageNum() int64 {
    return r._pageNum
}
// PageSize Setter
// 每页的数量。
func (r *AlibabaKbalgoAlscpoisGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaKbalgoAlscpoisGetRequest) GetPageSize() int64 {
    return r._pageSize
}