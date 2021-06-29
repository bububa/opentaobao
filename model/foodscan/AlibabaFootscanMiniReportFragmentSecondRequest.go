package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第二只脚生成报告接口 APIRequest
alibaba.footscan.mini.report.fragment.second

第二只脚生成报告接口
*/
type AlibabaFootscanMiniReportFragmentSecondRequest struct {
    model.Params

    // 平台分配的token
    token   string 

    // 请求数据
    reqData   *FilePackageBasicReq 

}

func NewAlibabaFootscanMiniReportFragmentSecondRequest() *AlibabaFootscanMiniReportFragmentSecondRequest{
    return &AlibabaFootscanMiniReportFragmentSecondRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.report.fragment.second"
}

func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFootscanMiniReportFragmentSecondRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetToken() string {
    return r.token
}

func (r *AlibabaFootscanMiniReportFragmentSecondRequest) SetReqData(reqData *FilePackageBasicReq) error {
    r.reqData = reqData
    r.Set("req_data", reqData)
    return nil
}

func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetReqData() *FilePackageBasicReq {
    return r.reqData
}

