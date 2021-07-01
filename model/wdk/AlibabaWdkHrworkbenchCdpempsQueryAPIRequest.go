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
type AlibabaWdkHrworkbenchCdpempsQueryAPIRequest struct {
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

// 初始化AlibabaWdkHrworkbenchCdpempsQueryAPIRequest对象
func NewAlibabaWdkHrworkbenchCdpempsQueryRequest() *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest{
    return &AlibabaWdkHrworkbenchCdpempsQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.hrworkbench.cdpemps.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 页面大小
func (r *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// BizKey Setter
// 业务授权key
func (r *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) SetBizKey(_bizKey string) error {
    r._bizKey = _bizKey
    r.Set("biz_key", _bizKey)
    return nil
}

// BizKey Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetBizKey() string {
    return r._bizKey
}
// BizCode Setter
// 业务授权code
func (r *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetBizCode() string {
    return r._bizCode
}
// CurrentPage Setter
// 起始页
func (r *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
