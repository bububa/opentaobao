package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第一只脚生成报告接口 API请求
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

// 初始化AlibabaFootscanMiniReportFragmentFirstRequest对象
func NewAlibabaFootscanMiniReportFragmentFirstRequest() *AlibabaFootscanMiniReportFragmentFirstRequest{
    return &AlibabaFootscanMiniReportFragmentFirstRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.report.fragment.first"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniReportFragmentFirstRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetToken() string {
    return r.token
}
// ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportFragmentFirstRequest) SetReqData(reqData *FilePackageRequest) error {
    r.reqData = reqData
    r.Set("req_data", reqData)
    return nil
}

// ReqData Getter
func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetReqData() *FilePackageRequest {
    return r.reqData
}
