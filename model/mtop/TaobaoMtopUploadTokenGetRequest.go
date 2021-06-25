package mtop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文件上传授权 APIRequest
taobao.mtop.upload.token.get

获取mtop文件上传授权
*/
type TaobaoMtopUploadTokenGetRequest struct {
    model.Params

    // 系统自动生成
    paramUploadTokenRequest   *UploadTokenRequestV 

}

func NewTaobaoMtopUploadTokenGetRequest() *TaobaoMtopUploadTokenGetRequest{
    return &TaobaoMtopUploadTokenGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMtopUploadTokenGetRequest) GetApiMethodName() string {
    return "taobao.mtop.upload.token.get"
}

func (r TaobaoMtopUploadTokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMtopUploadTokenGetRequest) SetParamUploadTokenRequest(paramUploadTokenRequest *UploadTokenRequestV) error {
    r.paramUploadTokenRequest = paramUploadTokenRequest
    r.Set("param_upload_token_request", paramUploadTokenRequest)
    return nil
}

func (r TaobaoMtopUploadTokenGetRequest) GetParamUploadTokenRequest() *UploadTokenRequestV {
    return r.paramUploadTokenRequest
}

