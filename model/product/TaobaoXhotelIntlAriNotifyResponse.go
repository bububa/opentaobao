package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
国际酒店集团价库变更通知 APIResponse
taobao.xhotel.intl.ari.notify

国际酒店集团价库变更时通知变更内容，平台及时更新价库信息，保证价库新鲜度
*/
type TaobaoXhotelIntlAriNotifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoXhotelIntlAriNotifyResponse `json:"xhotel_intl_ari_notify_response,omitempty"` 
    TaobaoXhotelIntlAriNotifyResponse
}

/* model for simplify = false
type TaobaoXhotelIntlAriNotifyResponse struct {

    // 通知结果
    
    Module  *struct {
        CacheChangeNotifyResult  *CacheChangeNotifyResult `json:"cache_change_notify_result,omitempty"`
    } `json:"module,omitempty"`
    

}
*/

type TaobaoXhotelIntlAriNotifyResponse struct {

    // 通知结果
    Module   *CacheChangeNotifyResult `json:"module,omitempty"`

}
