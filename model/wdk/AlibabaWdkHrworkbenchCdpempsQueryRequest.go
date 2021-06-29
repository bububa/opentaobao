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
    pageSize   int64
    // 业务授权key
    bizKey   string
    // 业务授权code
    bizCode   string
    // 起始页
    currentPage   int64
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
func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// BizKey Setter
// 业务授权key
func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetBizKey(bizKey string) error {
    r.bizKey = bizKey
    r.Set("biz_key", bizKey)
    return nil
}

// BizKey Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetBizKey() string {
    return r.bizKey
}
// BizCode Setter
// 业务授权code
func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

// BizCode Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetBizCode() string {
    return r.bizCode
}
// CurrentPage Setter
// 起始页
func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}
