package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询报告列表 API请求
alibaba.footscan.mini.report.list

查询报告列表
*/
type AlibabaFootscanMiniReportListRequest struct {
    model.Params
    // 平台分配的token
    token   string
    // 请求数据
    reqData   *TobFeetModelMobileReportRequest
}

// 初始化AlibabaFootscanMiniReportListRequest对象
func NewAlibabaFootscanMiniReportListRequest() *AlibabaFootscanMiniReportListRequest{
    return &AlibabaFootscanMiniReportListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniReportListRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.report.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniReportListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniReportListRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaFootscanMiniReportListRequest) GetToken() string {
    return r.token
}
// ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportListRequest) SetReqData(reqData *TobFeetModelMobileReportRequest) error {
    r.reqData = reqData
    r.Set("req_data", reqData)
    return nil
}

// ReqData Getter
func (r AlibabaFootscanMiniReportListRequest) GetReqData() *TobFeetModelMobileReportRequest {
    return r.reqData
}
