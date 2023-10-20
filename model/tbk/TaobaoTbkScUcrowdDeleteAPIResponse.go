package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscucrowddeleteAPIResponse 淘宝客-服务商-删除人群标签 API返回值
// taobao.tbk.sc.ucrowd.delete
//
// 服务商使用。支持淘宝客删除人群标签id，被删除的人群标签id将失效，并不可重新生效。
type TaobaotbkscucrowddeleteAPIResponse struct {
	model.CommonResponse
	TaobaotbkscucrowddeleteAPIResponseModel
}

// TaobaotbkscucrowddeleteAPIResponseModel is 淘宝客-服务商-删除人群标签 成功返回结果
type TaobaotbkscucrowddeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_ucrowd_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaotbkscucrowddeleteRpcResult `json:"data,omitempty" xml:"data,omitempty"`
}
