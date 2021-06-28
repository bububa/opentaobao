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
	RequestId     string         `json:"request_id,omitempty" xml:"files_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // results
    
    Results   []TopDownloadRecordDo `json:"results,omitempty" xml:"