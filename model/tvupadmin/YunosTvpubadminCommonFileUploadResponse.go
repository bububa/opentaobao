package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
文件上传API APIResponse
yunos.tvpubadmin.common.file.upload

文件上传服务
*/
type YunosTvpubadminCommonFileUploadAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminCommonFileUploadResponse
}

type YunosTvpubadminCommonFileUploadResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_common_file_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 上传后的文件地址
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
