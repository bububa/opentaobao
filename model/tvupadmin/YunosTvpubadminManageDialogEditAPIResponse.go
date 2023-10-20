package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageDialogEditAPIResponse 编辑全局弹窗 API返回值
// yunos.tvpubadmin.manage.dialog.edit
//
// 编辑全局弹窗
type YunosTvpubadminManageDialogEditAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageDialogEditAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageDialogEditAPIResponseModel).Reset()
}

// YunosTvpubadminManageDialogEditAPIResponseModel is 编辑全局弹窗 成功返回结果
type YunosTvpubadminManageDialogEditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvpubadminManageDialogEditTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminManageDialogEditAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageDialogEditAPIResponse)
	},
}

// GetYunosTvpubadminManageDialogEditAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageDialogEditAPIResponse
func GetYunosTvpubadminManageDialogEditAPIResponse() *YunosTvpubadminManageDialogEditAPIResponse {
	return poolYunosTvpubadminManageDialogEditAPIResponse.Get().(*YunosTvpubadminManageDialogEditAPIResponse)
}

// ReleaseYunosTvpubadminManageDialogEditAPIResponse 将 YunosTvpubadminManageDialogEditAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageDialogEditAPIResponse(v *YunosTvpubadminManageDialogEditAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageDialogEditAPIResponse.Put(v)
}
