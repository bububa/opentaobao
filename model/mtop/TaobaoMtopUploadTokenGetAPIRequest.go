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
type TaobaoMtopUploadTokenGetAPIRequest struct {
    model.Params
    // 系统自动生成
    _paramUploadTokenRequest   *UploadTokenRequestV
}

// 初始化TaobaoMtopUploadTokenGetAPIRequest对象
func NewTaobaoMtopUploadTokenGetRequest() *TaobaoMtopUploadTokenGetAPIRequest{
    return &TaobaoMtopUploadTokenGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMtopUploadTokenGetAPIRequest) GetApiMethodName() string {
    return "taobao.mtop.upload.token.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMtopUploadTokenGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamUploadTokenRequest Setter
// 系统自动生成
func (r *TaobaoMtopUploadTokenGetAPIRequest) SetParamUploadTokenRequest(_paramUploadTokenRequest *UploadTokenRequestV) error {
    r._paramUploadTokenRequest = _paramUploadTokenRequest
    r.Set("param_upload_token_request", _paramUploadTokenRequest)
    return nil
}

// ParamUploadTokenRequest Getter
func (r TaobaoMtopUploadTokenGetAPIRequest) GetParamUploadTokenRequest() *UploadTokenRequestV {
    return r._paramUploadTokenRequest
}
