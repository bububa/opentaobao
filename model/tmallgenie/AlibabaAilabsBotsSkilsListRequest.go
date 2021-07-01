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
type AlibabaAilabsBotsSkilsListAPIRequest struct {
    model.Params
    // 当前页
    _pageIndex   int64
    // 分页单位
    _pageSize   int64
}

// 初始化AlibabaAilabsBotsSkilsListAPIRequest对象
func NewAlibabaAilabsBotsSkilsListRequest() *AlibabaAilabsBotsSkilsListAPIRequest{
    return &AlibabaAilabsBotsSkilsListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsBotsSkilsListAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.bots.skils.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsBotsSkilsListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageIndex Setter
// 当前页
func (r *AlibabaAilabsBotsSkilsListAPIRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r AlibabaAilabsBotsSkilsListAPIRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 分页单位
func (r *AlibabaAilabsBotsSkilsListAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAilabsBotsSkilsListAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
