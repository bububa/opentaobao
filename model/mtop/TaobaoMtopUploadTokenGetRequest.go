package mtop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文件上传授权 API请求
taobao.mtop.upload.token.get

获取mtop文件上传授权
*/
type TaobaoMtopUploadTokenGetRequest struct {
    model.Params
    // 系统自动生成
    _paramUploadTokenRequest   *UploadTokenRequestV
}

// 初始化TaobaoMtopUploadTokenGetRequest对象
func NewTaobaoMtopUploadTokenGetRequest() *TaobaoMtopUploadTokenGetRequest{
    return &TaobaoMtopUploadTokenGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMtopUploadTokenGetRequest) GetApiMethodName() string {
    return "taobao.mtop.upload.token.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMtopUploadTokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamUploadTokenRequest Setter
// 系统自动生成
func (r *TaobaoMtopUploadTokenGetRequest) SetParamUploadTokenRequest(_paramUploadTokenRequest *UploadTokenRequestV) error {
    r._paramUploadTokenRequest = _paramUploadTokenRequest
    r.Set("param_upload_token_request", _paramUploadTokenRequest)
    return nil
}

// ParamUploadTokenRequest Getter
func (r TaobaoMtopUploadTokenGetRequest) GetParamUploadTokenRequest() *UploadTokenRequestV {
    return r._paramUploadTokenRequest
}
