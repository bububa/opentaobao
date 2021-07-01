package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增停开服任务 API返回值 
yunos.tvpubadmin.diccontroltask.add

新增停开服任务
*/
type YunosTvpubadminDiccontroltaskAddAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDiccontroltaskAddAPIResponseModel
}

// 新增停开服任务 成功返回结果
type YunosTvpubadminDiccontroltaskAddAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // object
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`
}
