package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
业务文件获取 APIResponse
taobao.files.get

获取业务方暂存给ISV的文件列表
*/
type TaobaoFilesGetAPIResponse struct {
    model.CommonResponse
    TaobaoFilesGetResponse
}

type TaobaoFilesGetResponse struct {
    XMLName xml.Name `xml:"files_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // results
    
    Results   []TopDownloadRecordDo `json:"results,omitempty" xml:"results>top_download_record_do,omitempty"`
    
    
}
