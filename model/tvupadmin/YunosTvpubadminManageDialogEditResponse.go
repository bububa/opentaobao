package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑全局弹窗 APIResponse
yunos.tvpubadmin.manage.dialog.edit

编辑全局弹窗
*/
type YunosTvpubadminManageDialogEditAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageDialogEditResponse
}

type YunosTvpubadminManageDialogEditResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_edit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
