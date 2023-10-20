package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagedialogaddAPIResponse 新增全局弹窗 API返回值
// yunos.tvpubadmin.manage.dialog.add
//
// 新增全局弹窗
type YunostvpubadminmanagedialogaddAPIResponse struct {
	model.CommonResponse
	YunostvpubadminmanagedialogaddAPIResponseModel
}

// YunostvpubadminmanagedialogaddAPIResponseModel is 新增全局弹窗 成功返回结果
type YunostvpubadminmanagedialogaddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_dialog_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunostvpubadminmanagedialogaddTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
