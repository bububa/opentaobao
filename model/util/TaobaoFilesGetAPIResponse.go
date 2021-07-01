package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
业务文件获取 API返回值 
taobao.files.get

获取业务方暂存给ISV的文件列表
*/
type TaobaoFilesGetAPIResponse struct {
    model.CommonResponse
    TaobaoFilesGetAPIResponseModel
}

// 业务文件获取 成功返回结果
type TaobaoFilesGetAPIResponseModel struct {
    XMLName xml.Name `xml:"files_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // results
    Results   []TopDownloadRecordDo `json:"results,omitempty" xml:"results>top_download_record_do,omitempty"`
}
