package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
业务文件获取 APIResponse
taobao.files.get

获取业务方暂存给ISV的文件列表
*/
type TaobaoFilesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFilesGetResponse `json:"files_get_response,omitempty"` 
    TaobaoFilesGetResponse
}

/* model for simplify = false
type TaobaoFilesGetResponse struct {

    // results
    
    Results  struct {
        TopDownloadRecordDo  []TopDownloadRecordDo `json:"top_download_record_do,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoFilesGetResponse struct {

    // results
    Results   []TopDownloadRecordDo `json:"results,omitempty"`

}
