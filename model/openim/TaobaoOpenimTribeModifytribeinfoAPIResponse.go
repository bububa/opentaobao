package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeModifytribeinfoAPIResponse OPENIM群信息修改 API返回值
// taobao.openim.tribe.modifytribeinfo
//
// OPENIM群信息修改
type TaobaoOpenimTribeModifytribeinfoAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeModifytribeinfoAPIResponseModel
}

// TaobaoOpenimTribeModifytribeinfoAPIResponseModel is OPENIM群信息修改 成功返回结果
type TaobaoOpenimTribeModifytribeinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_modifytribeinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}
