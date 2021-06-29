package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
homs员工信息核对查询服务 API请求
alibaba.wdk.hrworkbench.cdpemps.query

给盒马可靠软件服务商Cdp系统，做非阿里编员工数据一致性核对检查
*/
type AlibabaWdkHrworkbenchCdpempsQueryRequest struct {
    model.Params
    // 页面大小
    _pageSize   int64
    // 业务授权key
    _bizKey   string
    // 业务授权code
    _bizCode   string
    // 起始页
    _currentPage   int64
}

// 初始化AlibabaWdkHrworkbenchCdpempsQueryRequest对象
func NewAlibabaWdkHrworkbenchCdpempsQueryRequest() *AlibabaWdkHrworkbenchCdpempsQueryRequest{
    return &AlibabaWdkHrworkbenchCdpempsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.hrworkbench.cdpemps.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 页面大小
func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// BizKey Setter
// 业务授权key
func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetBizKey(_bizKey string) error {
    r._bizKey = _bizKey
    r.Set("biz_key", _bizKey)
    return nil
}

// BizKey Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetBizKey() string {
    return r._bizKey
}
// BizCode Setter
// 业务授权code
func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetBizCode() string {
    return r._bizCode
}
// CurrentPage Setter
// 起始页
func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetCurrentPage() int64 {
    return r._currentPage
}
