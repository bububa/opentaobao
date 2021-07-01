package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageDialogAddAPIResponse
新增全局弹窗 API返回值
yunos.tvpubadmin.manage.dialog.add

新增全局弹窗 */
type YunosTvpubadminManageDialogAddAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageDialogAddAPIResponseModel
}

// YunosTvpubadminManageDialogAddAPIResponseModel is 新增全局弹窗 成功返回结果
type YunosTvpubadminManageDialogAddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResult `json:"result,omitempty" xml:"result,omitempty"`
}
