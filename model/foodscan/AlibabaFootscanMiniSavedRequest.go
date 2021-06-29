package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新报告状态 APIRequest
alibaba.footscan.mini.saved

更新报告状态接口
*/
type AlibabaFootscanMiniSavedRequest struct {
    model.Params

    // 平台分配的token
    token   string 

    // 请求数据
    reqData   string 

}

func NewAlibabaFootscanMiniSavedRequest() *AlibabaFootscanMiniSavedRequest{
    return &AlibabaFootscanMiniSavedRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFootscanMiniSavedRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.saved"
}

func (r AlibabaFootscanMiniSavedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFootscanMiniSavedRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaFootscanMiniSavedRequest) GetToken() string {
    return r.token
}

func (r *AlibabaFootscanMiniSavedRequest) SetReqData(reqData string) error {
    r.reqData = reqData
    r.Set("req_data", reqData)
    return nil
}

func (r AlibabaFootscanMiniSavedRequest) GetReqData() string {
    return r.reqData
}

