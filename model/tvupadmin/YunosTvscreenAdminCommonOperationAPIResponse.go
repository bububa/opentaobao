package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvscreenAdminCommonOperationAPIResponse 一体机桌面通用接口 API返回值
// yunos.tvscreen.admin.common.operation
//
// 一体机桌面通用接口
type YunosTvscreenAdminCommonOperationAPIResponse struct {
	model.CommonResponse
	YunosTvscreenAdminCommonOperationAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvscreenAdminCommonOperationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvscreenAdminCommonOperationAPIResponseModel).Reset()
}

// YunosTvscreenAdminCommonOperationAPIResponseModel is 一体机桌面通用接口 成功返回结果
type YunosTvscreenAdminCommonOperationAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvscreen_admin_common_operation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvscreenAdminCommonOperationTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvscreenAdminCommonOperationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvscreenAdminCommonOperationAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvscreenAdminCommonOperationAPIResponse)
	},
}

// GetYunosTvscreenAdminCommonOperationAPIResponse 从 sync.Pool 获取 YunosTvscreenAdminCommonOperationAPIResponse
func GetYunosTvscreenAdminCommonOperationAPIResponse() *YunosTvscreenAdminCommonOperationAPIResponse {
	return poolYunosTvscreenAdminCommonOperationAPIResponse.Get().(*YunosTvscreenAdminCommonOperationAPIResponse)
}

// ReleaseYunosTvscreenAdminCommonOperationAPIResponse 将 YunosTvscreenAdminCommonOperationAPIResponse 保存到 sync.Pool
func ReleaseYunosTvscreenAdminCommonOperationAPIResponse(v *YunosTvscreenAdminCommonOperationAPIResponse) {
	v.Reset()
	poolYunosTvscreenAdminCommonOperationAPIResponse.Put(v)
}
