package kbalgo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百度批量获取本地poi接口 APIRequest
alibaba.kbalgo.alscpois.get

接口用于百度方获取本地生活poi数据，分页获取。
*/
type AlibabaKbalgoAlscpoisGetRequest struct {
    model.Params

    // 第几页
    pageNum   int64 

    // 每页的数量。
    pageSize   int64 

}

func NewAlibabaKbalgoAlscpoisGetRequest() *AlibabaKbalgoAlscpoisGetRequest{
    return &AlibabaKbalgoAlscpoisGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaKbalgoAlscpoisGetRequest) GetApiMethodName() string {
    return "alibaba.kbalgo.alscpois.get"
}

func (r AlibabaKbalgoAlscpoisGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaKbalgoAlscpoisGetRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r AlibabaKbalgoAlscpoisGetRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *AlibabaKbalgoAlscpoisGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaKbalgoAlscpoisGetRequest) GetPageSize() int64 {
    return r.pageSize
}

