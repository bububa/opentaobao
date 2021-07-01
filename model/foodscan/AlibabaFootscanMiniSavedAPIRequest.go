package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新报告状态 API请求
alibaba.footscan.mini.saved

更新报告状态接口
*/
type AlibabaFootscanMiniSavedAPIRequest struct {
    model.Params
    // 平台分配的token
    _token   string
    // 请求数据
    _reqData   string
}

// 初始化AlibabaFootscanMiniSavedAPIRequest对象
func NewAlibabaFootscanMiniSavedRequest() *AlibabaFootscanMiniSavedAPIRequest{
    return &AlibabaFootscanMiniSavedAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniSavedAPIRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.saved"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniSavedAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniSavedAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaFootscanMiniSavedAPIRequest) GetToken() string {
    return r._token
}
// ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniSavedAPIRequest) SetReqData(_reqData string) error {
    r._reqData = _reqData
    r.Set("req_data", _reqData)
    return nil
}

// ReqData Getter
func (r AlibabaFootscanMiniSavedAPIRequest) GetReqData() string {
    return r._reqData
}
