package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmMenuListAPIResponse
获取特价菜单 API返回值
alibaba.alsc.crm.menu.list

获取特价菜单 */
type AlibabaAlscCrmMenuListAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmMenuListAPIResponseModel
}

// AlibabaAlscCrmMenuListAPIResponseModel is 获取特价菜单 成功返回结果
type AlibabaAlscCrmMenuListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_menu_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
