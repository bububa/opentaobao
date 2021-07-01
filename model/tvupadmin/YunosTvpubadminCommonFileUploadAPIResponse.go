package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
文件上传API API返回值 
yunos.tvpubadmin.common.file.upload

文件上传服务
*/
type YunosTvpubadminCommonFileUploadAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminCommonFileUploadAPIResponseModel
}

// 文件上传API 成功返回结果
type YunosTvpubadminCommonFileUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_common_file_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 上传后的文件地址
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
