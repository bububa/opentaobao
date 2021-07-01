package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvscreenAdminCommonOperationAPIResponse
一体机桌面通用接口 API返回值
yunos.tvscreen.admin.common.operation

一体机桌面通用接口 */
type YunosTvscreenAdminCommonOperationAPIResponse struct {
	model.CommonResponse
	YunosTvscreenAdminCommonOperationAPIResponseModel
}

// YunosTvscreenAdminCommonOperationAPIResponseModel is 一体机桌面通用接口 成功返回结果
type YunosTvscreenAdminCommonOperationAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvscreen_admin_common_operation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvscreenAdminCommonOperationTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
