package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScUcrowdGetAPIResponse 淘宝客-服务商-获取人群标签 API返回值
// taobao.tbk.sc.ucrowd.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群信息，包括人群名称描述及覆盖会员数。
type TaobaoTbkScUcrowdGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScUcrowdGetAPIResponseModel
}

// TaobaoTbkScUcrowdGetAPIResponseModel is 淘宝客-服务商-获取人群标签 成功返回结果
type TaobaoTbkScUcrowdGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_ucrowd_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScUcrowdGetRpcResult `json:"data,omitempty" xml:"data,omitempty"`
}
