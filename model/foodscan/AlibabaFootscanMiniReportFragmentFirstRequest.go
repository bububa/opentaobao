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
    _token   string
    // 请求数据
    _reqData   *FilePackageRequest
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
func (r *AlibabaFootscanMiniReportFragmentFirstRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetToken() string {
    return r._token
}
// ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportFragmentFirstRequest) SetReqData(_reqData *FilePackageRequest) error {
    r._reqData = _reqData
    r.Set("req_data", _reqData)
    return nil
}

// ReqData Getter
func (r AlibabaFootscanMiniReportFragmentFirstRequest) GetReqData() *FilePackageRequest {
    return r._reqData
}
