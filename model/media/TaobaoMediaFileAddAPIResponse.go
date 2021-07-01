package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
多媒体平台文件添加 API返回值 
taobao.media.file.add

用户通过top上传文件到多媒体平台
*/
type TaobaoMediaFileAddAPIResponse struct {
    model.CommonResponse
    TaobaoMediaFileAddAPIResponseModel
}

// 多媒体平台文件添加 成功返回结果
type TaobaoMediaFileAddAPIResponseModel struct {
    XMLName xml.Name `xml:"media_file_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 上传到多媒体平台的文件
    File   *File `json:"file,omitempty" xml:"file,omitempty"`
}
