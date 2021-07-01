package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
停开服任务状态变更 API返回值 
yunos.tvpubadmin.diccontroltask.update

停开服任务状态变更
*/
type YunosTvpubadminDiccontroltaskUpdateAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDiccontroltaskUpdateAPIResponseModel
}

// 停开服任务状态变更 成功返回结果
type YunosTvpubadminDiccontroltaskUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // object
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`
}
