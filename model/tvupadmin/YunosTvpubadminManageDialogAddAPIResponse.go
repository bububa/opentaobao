package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageDialogAddAPIResponse 新增全局弹窗 API返回值
// yunos.tvpubadmin.manage.dialog.add
//
// 新增全局弹窗
type YunosTvpubadminManageDialogAddAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageDialogAddAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageDialogAddAPIResponseModel).Reset()
}

// YunosTvpubadminManageDialogAddAPIResponseModel is 新增全局弹窗 成功返回结果
type YunosTvpubadminManageDialogAddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvpubadminManageDialogAddTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminManageDialogAddAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageDialogAddAPIResponse)
	},
}

// GetYunosTvpubadminManageDialogAddAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageDialogAddAPIResponse
func GetYunosTvpubadminManageDialogAddAPIResponse() *YunosTvpubadminManageDialogAddAPIResponse {
	return poolYunosTvpubadminManageDialogAddAPIResponse.Get().(*YunosTvpubadminManageDialogAddAPIResponse)
}

// ReleaseYunosTvpubadminManageDialogAddAPIResponse 将 YunosTvpubadminManageDialogAddAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageDialogAddAPIResponse(v *YunosTvpubadminManageDialogAddAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageDialogAddAPIResponse.Put(v)
}
