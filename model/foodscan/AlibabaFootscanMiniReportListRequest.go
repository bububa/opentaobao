package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询报告列表 APIRequest
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

func NewAlibabaFootscanMiniReportListRequest() *AlibabaFootscanMiniReportListRequest{
    return &AlibabaFootscanMiniReportListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFootscanMiniReportListRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.report.list"
}

func (r AlibabaFootscanMiniReportListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFootscanMiniReportListRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaFootscanMiniReportListRequest) GetToken() string {
    return r.token
}

func (r *AlibabaFootscanMiniReportListRequest) SetReqData(reqData *TobFeetModelMobileReportRequest) error {
    r.reqData = reqData
    r.Set("req_data", reqData)
    return nil
}

func (r AlibabaFootscanMiniReportListRequest) GetReqData() *TobFeetModelMobileReportRequest {
    return r.reqData
}

