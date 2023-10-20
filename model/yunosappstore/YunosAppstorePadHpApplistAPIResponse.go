package yunosappstore

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosAppstorePadHpApplistAPIResponse 查询HpPad appList API返回值
// yunos.appstore.pad.hp.applist
//
// 提供hp pad应用群数据
type YunosAppstorePadHpApplistAPIResponse struct {
	model.CommonResponse
	YunosAppstorePadHpApplistAPIResponseModel
}

// YunosAppstorePadHpApplistAPIResponseModel is 查询HpPad appList 成功返回结果
type YunosAppstorePadHpApplistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_appstore_pad_hp_applist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []YunosAppstorePadHpApplistResult `json:"results,omitempty" xml:"results>yunos_appstore_pad_hp_applist_result,omitempty"`
}
