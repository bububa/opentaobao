package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第一只脚生成报告接口 APIRequest
alibaba.footscan.mini.report.fragment.first

第一只脚生成报告接口
*/
type AlibabaFootscanMiniReportFragmentFirstRequest struct {
    model.Params

    // 平台分配的token
    token   string 

    // 请求数据
    reqData   *FilePackageRequest 

}

func NewAlibabaFootscanMiniReportFragmentFirstRequest() *AlibabaFootscanMiniReportFragmentFirstRequest{
    return &AlibabaFootscanMiniReportFragmentFirstRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.report.fragment.first"
}

func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFootscanMiniReportFragmentFirstRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetToken() string {
    return r.token
}

func (r *AlibabaFootscanMiniReportFragmentFirstRequest) SetReqData(reqData *FilePackageRequest) error {
    r.reqData = reqData
    r.Set("req_data", reqData)
    return nil
}

func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetReqData() *FilePackageRequest {
    return r.reqData
}

