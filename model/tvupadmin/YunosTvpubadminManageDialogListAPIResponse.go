package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageDialogListAPIResponse 分页获取弹窗列表 API返回值
// yunos.tvpubadmin.manage.dialog.list
//
// 分页获取弹窗配置列表
type YunosTvpubadminManageDialogListAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageDialogListAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageDialogListAPIResponseModel).Reset()
}

// YunosTvpubadminManageDialogListAPIResponseModel is 分页获取弹窗列表 成功返回结果
type YunosTvpubadminManageDialogListAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminManageDialogListAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageDialogListAPIResponse)
	},
}

// GetYunosTvpubadminManageDialogListAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageDialogListAPIResponse
func GetYunosTvpubadminManageDialogListAPIResponse() *YunosTvpubadminManageDialogListAPIResponse {
	return poolYunosTvpubadminManageDialogListAPIResponse.Get().(*YunosTvpubadminManageDialogListAPIResponse)
}

// ReleaseYunosTvpubadminManageDialogListAPIResponse 将 YunosTvpubadminManageDialogListAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageDialogListAPIResponse(v *YunosTvpubadminManageDialogListAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageDialogListAPIResponse.Put(v)
}
