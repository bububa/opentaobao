package mtop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文件上传授权 APIResponse
taobao.mtop.upload.token.get

获取mtop文件上传授权
*/
type TaobaoMtopUploadTokenGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"mtop_upload_token_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // msg
    
    Message   string `json:"message,omitempty" xml:"