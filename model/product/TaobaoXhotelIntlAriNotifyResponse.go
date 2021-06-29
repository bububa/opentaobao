package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际酒店集团价库变更通知 APIResponse
taobao.xhotel.intl.ari.notify

国际酒店集团价库变更时通知变更内容，平台及时更新价库信息，保证价库新鲜度
*/
type TaobaoXhotelIntlAriNotifyAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelIntlAriNotifyResponse
}

type TaobaoXhotelIntlAriNotifyResponse struct {
    XMLName xml.Name `xml:"xhotel_intl_ari_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 通知结果
    
    Module   *CacheChangeNotifyResult `json:"module,omitempty" xml:"module,omitempty"`

    
}
