package yunosappstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询HpPad appList APIResponse
yunos.appstore.pad.hp.applist

提供hp pad应用群数据
*/
type YunosAppstorePadHpApplistAPIResponse struct {
    model.CommonResponse
    YunosAppstorePadHpApplistResponse
}

type YunosAppstorePadHpApplistResponse struct {
    XMLName xml.Name `xml:"yunos_appstore_pad_hp_applist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Results   []YunosAppstorePadHpApplistResult `json:"results,omitempty" xml:"results>yunos_appstore_pad_hp_applist_result,omitempty"`
    
    
}
