package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据scanId查询报告 API请求
alibaba.footscan.mini.query.mobilereport

根据scanId查询报告
*/
type AlibabaFootscanMiniQueryMobilereportRequest struct {
    model.Params
    // 平台分配的token
    _token   string
    // 扫描ID
    _scanId   string
}

// 初始化AlibabaFootscanMiniQueryMobilereportRequest对象
func NewAlibabaFootscanMiniQueryMobilereportRequest() *AlibabaFootscanMiniQueryMobilereportRequest{
    return &AlibabaFootscanMiniQueryMobilereportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniQueryMobilereportRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.query.mobilereport"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniQueryMobilereportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniQueryMobilereportRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaFootscanMiniQueryMobilereportRequest) GetToken() string {
    return r._token
}
// ScanId Setter
// 扫描ID
func (r *AlibabaFootscanMiniQueryMobilereportRequest) SetScanId(_scanId string) error {
    r._scanId = _scanId
    r.Set("scan_id", _scanId)
    return nil
}

// ScanId Getter
func (r AlibabaFootscanMiniQueryMobilereportRequest) GetScanId() string {
    return r._scanId
}
