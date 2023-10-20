package yunosappstore

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosappstorepadhpapplistAPIResponse 查询HpPad appList API返回值
// yunos.appstore.pad.hp.applist
//
// 提供hp pad应用群数据
type YunosappstorepadhpapplistAPIResponse struct {
	model.CommonResponse
	YunosappstorepadhpapplistAPIResponseModel
}

// YunosappstorepadhpapplistAPIResponseModel is 查询HpPad appList 成功返回结果
type YunosappstorepadhpapplistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_appstore_pad_hp_applist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []YunosappstorepadhpapplistResult `json:"results,omitempty" xml:"results>yunosappstorepadhpapplist_result,omitempty"`
}
