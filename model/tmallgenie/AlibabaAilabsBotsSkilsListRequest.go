package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对外设备获取技能列表 APIRequest
alibaba.ailabs.bots.skils.list

获取ai开放平台技能列表
*/
type AlibabaAilabsBotsSkilsListRequest struct {
    model.Params

    // 当前页
    pageIndex   int64 

    // 分页单位
    pageSize   int64 

}

func NewAlibabaAilabsBotsSkilsListRequest() *AlibabaAilabsBotsSkilsListRequest{
    return &AlibabaAilabsBotsSkilsListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsBotsSkilsListRequest) GetApiMethodName() string {
    return "alibaba.ailabs.bots.skils.list"
}

func (r AlibabaAilabsBotsSkilsListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsBotsSkilsListRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r AlibabaAilabsBotsSkilsListRequest) GetPageIndex() int64 {
    return r.pageIndex
}

func (r *AlibabaAilabsBotsSkilsListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAilabsBotsSkilsListRequest) GetPageSize() int64 {
    return r.pageSize
}

