package nazca

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文件秘钥 APIResponse
alibaba.nazca.token.filesecret.get

获取文件秘钥
*/
type AlibabaNazcaTokenFilesecretGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_nazca_token_filesecret_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // error
    
    Error   string `json:"error,omitempty" xml:"