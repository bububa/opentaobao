package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeCreateAPIResponse 创建群 API返回值
// taobao.openim.tribe.create
//
// 创建一个openim的群
type TaobaoOpenimTribeCreateAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeCreateAPIResponseModel
}

// TaobaoOpenimTribeCreateAPIResponseModel is 创建群 成功返回结果
type TaobaoOpenimTribeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建群的信息
	TribeInfo *TribeInfo `json:"tribe_info,omitempty" xml:"tribe_info,omitempty"`
}
