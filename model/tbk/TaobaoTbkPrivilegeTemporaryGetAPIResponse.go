package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkPrivilegeTemporaryGetAPIResponse 淘宝客-服务商-单品券高效转链（临时接口） API返回值
// taobao.tbk.privilege.temporary.get
//
// 单品券高效转链API
type TaobaoTbkPrivilegeTemporaryGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkPrivilegeTemporaryGetAPIResponseModel
}

// TaobaoTbkPrivilegeTemporaryGetAPIResponseModel is 淘宝客-服务商-单品券高效转链（临时接口） 成功返回结果
type TaobaoTbkPrivilegeTemporaryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_privilege_temporary_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTbkPrivilegeTemporaryGetRpcResult `json:"result,omitempty" xml:"result,omitempty"`
}
