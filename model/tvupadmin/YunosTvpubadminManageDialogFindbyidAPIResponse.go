package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id查询全局弹窗 API返回值 
yunos.tvpubadmin.manage.dialog.findbyid

根据id查询全局弹窗
*/
type YunosTvpubadminManageDialogFindbyidAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageDialogFindbyidAPIResponseModel
}

// 根据id查询全局弹窗 成功返回结果
type YunosTvpubadminManageDialogFindbyidAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_findbyid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // object
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
