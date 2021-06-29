package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第二只脚生成报告接口 API请求
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

// 初始化AlibabaFootscanMiniReportFragmentSecondRequest对象
func NewAlibabaFootscanMiniReportFragmentSecondRequest() *AlibabaFootscanMiniReportFragmentSecondRequest{
    return &AlibabaFootscanMiniReportFragmentSecondRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.report.fragment.second"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniReportFragmentSecondRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetToken() string {
    return r.token
}
// ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportFragmentSecondRequest) SetReqData(reqData *FilePackageBasicReq) error {
    r.reqData = reqData
    r.Set("req_data", reqData)
    return nil
}

// ReqData Getter
func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetReqData() *FilePackageBasicReq {
    return r.reqData
}
