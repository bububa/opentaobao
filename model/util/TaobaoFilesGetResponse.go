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
    Response *TaobaoFilesGetResponse `json:"taobao_files_get_response,omitempty"`
}

type TaobaoFilesGetResponse struct {

    // results
    Results   []TopDownloadRecordDo `json:"results,omitempty"`

}
