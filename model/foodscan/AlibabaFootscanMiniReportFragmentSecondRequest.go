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
    _token   string
    // 请求数据
    _reqData   *FilePackageBasicReq
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
func (r *AlibabaFootscanMiniReportFragmentSecondRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetToken() string {
    return r._token
}
// ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportFragmentSecondRequest) SetReqData(_reqData *FilePackageBasicReq) error {
    r._reqData = _reqData
    r.Set("req_data", _reqData)
    return nil
}

// ReqData Getter
func (r AlibabaFootscanMiniReportFragmentSecondRequest) GetReqData() *FilePackageBasicReq {
    return r._reqData
}
