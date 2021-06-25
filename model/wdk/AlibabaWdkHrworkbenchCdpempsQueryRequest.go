package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
homs员工信息核对查询服务 APIRequest
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

func NewAlibabaWdkHrworkbenchCdpempsQueryRequest() *AlibabaWdkHrworkbenchCdpempsQueryRequest{
    return &AlibabaWdkHrworkbenchCdpempsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.hrworkbench.cdpemps.query"
}

func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetBizKey(bizKey string) error {
    r.bizKey = bizKey
    r.Set("biz_key", bizKey)
    return nil
}

func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetBizKey() string {
    return r.bizKey
}

func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetBizCode() string {
    return r.bizCode
}

func (r *AlibabaWdkHrworkbenchCdpempsQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaWdkHrworkbenchCdpempsQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}

