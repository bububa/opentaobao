package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageDialogDeleteAPIResponse 删除全局弹窗 API返回值
// yunos.tvpubadmin.manage.dialog.delete
//
// 删除全局弹窗
type YunosTvpubadminManageDialogDeleteAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageDialogDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageDialogDeleteAPIResponseModel).Reset()
}

// YunosTvpubadminManageDialogDeleteAPIResponseModel is 删除全局弹窗 成功返回结果
type YunosTvpubadminManageDialogDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvpubadminManageDialogDeleteTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminManageDialogDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageDialogDeleteAPIResponse)
	},
}

// GetYunosTvpubadminManageDialogDeleteAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageDialogDeleteAPIResponse
func GetYunosTvpubadminManageDialogDeleteAPIResponse() *YunosTvpubadminManageDialogDeleteAPIResponse {
	return poolYunosTvpubadminManageDialogDeleteAPIResponse.Get().(*YunosTvpubadminManageDialogDeleteAPIResponse)
}

// ReleaseYunosTvpubadminManageDialogDeleteAPIResponse 将 YunosTvpubadminManageDialogDeleteAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageDialogDeleteAPIResponse(v *YunosTvpubadminManageDialogDeleteAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageDialogDeleteAPIResponse.Put(v)
}
