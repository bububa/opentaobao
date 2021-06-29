package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提供给外部系统的免登校验 APIRequest
alibaba.happytrip.freelogin.getusercontext

免登融合，提供免登相关接口给外部供应商做登录验证
*/
type AlibabaHappytripFreeloginGetusercontextRequest struct {
    model.Params

    // 请求入参
    req   *SsoParamDto 

}

func NewAlibabaHappytripFreeloginGetusercontextRequest() *AlibabaHappytripFreeloginGetusercontextRequest{
    return &AlibabaHappytripFreeloginGetusercontextRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripFreeloginGetusercontextRequest) GetApiMethodName() string {
    return "alibaba.happytrip.freelogin.getusercontext"
}

func (r AlibabaHappytripFreeloginGetusercontextRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripFreeloginGetusercontextRequest) SetReq(req *SsoParamDto) error {
    r.req = req
    r.Set("req", req)
    return nil
}

func (r AlibabaHappytripFreeloginGetusercontextRequest) GetReq() *SsoParamDto {
    return r.req
}

