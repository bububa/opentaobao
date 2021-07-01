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
type AlibabaHappytripFreeloginGetusercontextAPIRequest struct {
    model.Params
    // 请求入参
    _req   *SsoParamDto
}

// 初始化AlibabaHappytripFreeloginGetusercontextAPIRequest对象
func NewAlibabaHappytripFreeloginGetusercontextRequest() *AlibabaHappytripFreeloginGetusercontextAPIRequest{
    return &AlibabaHappytripFreeloginGetusercontextAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripFreeloginGetusercontextAPIRequest) GetApiMethodName() string {
    return "alibaba.happytrip.freelogin.getusercontext"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripFreeloginGetusercontextAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Req Setter
// 请求入参
func (r *AlibabaHappytripFreeloginGetusercontextAPIRequest) SetReq(_req *SsoParamDto) error {
    r._req = _req
    r.Set("req", _req)
    return nil
}

// Req Getter
func (r AlibabaHappytripFreeloginGetusercontextAPIRequest) GetReq() *SsoParamDto {
    return r._req
}
