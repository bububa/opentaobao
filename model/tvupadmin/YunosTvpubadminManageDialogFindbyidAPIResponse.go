package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageDialogFindbyidAPIResponse 根据id查询全局弹窗 API返回值
// yunos.tvpubadmin.manage.dialog.findbyid
//
// 根据id查询全局弹窗
type YunosTvpubadminManageDialogFindbyidAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageDialogFindbyidAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogFindbyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageDialogFindbyidAPIResponseModel).Reset()
}

// YunosTvpubadminManageDialogFindbyidAPIResponseModel is 根据id查询全局弹窗 成功返回结果
type YunosTvpubadminManageDialogFindbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_findbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageDialogFindbyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminManageDialogFindbyidAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageDialogFindbyidAPIResponse)
	},
}

// GetYunosTvpubadminManageDialogFindbyidAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageDialogFindbyidAPIResponse
func GetYunosTvpubadminManageDialogFindbyidAPIResponse() *YunosTvpubadminManageDialogFindbyidAPIResponse {
	return poolYunosTvpubadminManageDialogFindbyidAPIResponse.Get().(*YunosTvpubadminManageDialogFindbyidAPIResponse)
}

// ReleaseYunosTvpubadminManageDialogFindbyidAPIResponse 将 YunosTvpubadminManageDialogFindbyidAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageDialogFindbyidAPIResponse(v *YunosTvpubadminManageDialogFindbyidAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageDialogFindbyidAPIResponse.Put(v)
}
