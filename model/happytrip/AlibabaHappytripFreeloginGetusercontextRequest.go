package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提供给外部系统的免登校验 API请求
alibaba.happytrip.freelogin.getusercontext

免登融合，提供免登相关接口给外部供应商做登录验证
*/
type AlibabaHappytripFreeloginGetusercontextRequest struct {
    model.Params
    // 请求入参
    _req   *SsoParamDto
}

// 初始化AlibabaHappytripFreeloginGetusercontextRequest对象
func NewAlibabaHappytripFreeloginGetusercontextRequest() *AlibabaHappytripFreeloginGetusercontextRequest{
    return &AlibabaHappytripFreeloginGetusercontextRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripFreeloginGetusercontextRequest) GetApiMethodName() string {
    return "alibaba.happytrip.freelogin.getusercontext"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripFreeloginGetusercontextRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Req Setter
// 请求入参
func (r *AlibabaHappytripFreeloginGetusercontextRequest) SetReq(_req *SsoParamDto) error {
    r._req = _req
    r.Set("req", _req)
    return nil
}

// Req Getter
func (r AlibabaHappytripFreeloginGetusercontextRequest) GetReq() *SsoParamDto {
    return r._req
}
