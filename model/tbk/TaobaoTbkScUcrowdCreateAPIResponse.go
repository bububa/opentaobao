package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscucrowdcreateAPIResponse 淘宝客-服务商-创建人群标签 API返回值
// taobao.tbk.sc.ucrowd.create
//
// 服务商使用。可为淘宝客创建会员人群标签，获得人群id，人群id可用于物料评估推荐等场景。
type TaobaotbkscucrowdcreateAPIResponse struct {
	model.CommonResponse
	TaobaotbkscucrowdcreateAPIResponseModel
}

// TaobaotbkscucrowdcreateAPIResponseModel is 淘宝客-服务商-创建人群标签 成功返回结果
type TaobaotbkscucrowdcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_ucrowd_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaotbkscucrowdcreateRpcResult `json:"data,omitempty" xml:"data,omitempty"`
}
