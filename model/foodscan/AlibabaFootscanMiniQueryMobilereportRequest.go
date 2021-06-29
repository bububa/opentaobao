package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据scanId查询报告 APIRequest
alibaba.footscan.mini.query.mobilereport

根据scanId查询报告
*/
type AlibabaFootscanMiniQueryMobilereportRequest struct {
    model.Params

    // 平台分配的token
    token   string 

    // 扫描ID
    scanId   string 

}

func NewAlibabaFootscanMiniQueryMobilereportRequest() *AlibabaFootscanMiniQueryMobilereportRequest{
    return &AlibabaFootscanMiniQueryMobilereportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFootscanMiniQueryMobilereportRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.query.mobilereport"
}

func (r AlibabaFootscanMiniQueryMobilereportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFootscanMiniQueryMobilereportRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaFootscanMiniQueryMobilereportRequest) GetToken() string {
    return r.token
}

func (r *AlibabaFootscanMiniQueryMobilereportRequest) SetScanId(scanId string) error {
    r.scanId = scanId
    r.Set("scan_id", scanId)
    return nil
}

func (r AlibabaFootscanMiniQueryMobilereportRequest) GetScanId() string {
    return r.scanId
}

