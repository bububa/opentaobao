package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取弹窗列表 API返回值 
yunos.tvpubadmin.manage.dialog.list

分页获取弹窗配置列表
*/
type YunosTvpubadminManageDialogListAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageDialogListAPIResponseModel
}

// 分页获取弹窗列表 成功返回结果
type YunosTvpubadminManageDialogListAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // object
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
