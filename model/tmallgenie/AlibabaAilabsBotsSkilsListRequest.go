package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对外设备获取技能列表 API请求
alibaba.ailabs.bots.skils.list

获取ai开放平台技能列表
*/
type AlibabaAilabsBotsSkilsListRequest struct {
    model.Params
    // 当前页
    _pageIndex   int64
    // 分页单位
    _pageSize   int64
}

// 初始化AlibabaAilabsBotsSkilsListRequest对象
func NewAlibabaAilabsBotsSkilsListRequest() *AlibabaAilabsBotsSkilsListRequest{
    return &AlibabaAilabsBotsSkilsListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsBotsSkilsListRequest) GetApiMethodName() string {
    return "alibaba.ailabs.bots.skils.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsBotsSkilsListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageIndex Setter
// 当前页
func (r *AlibabaAilabsBotsSkilsListRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r AlibabaAilabsBotsSkilsListRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 分页单位
func (r *AlibabaAilabsBotsSkilsListRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAilabsBotsSkilsListRequest) GetPageSize() int64 {
    return r._pageSize
}
