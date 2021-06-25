package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际酒店集团价库变更通知 APIRequest
taobao.xhotel.intl.ari.notify

国际酒店集团价库变更时通知变更内容，平台及时更新价库信息，保证价库新鲜度
*/
type TaobaoXhotelIntlAriNotifyRequest struct {
    model.Params

    // 缓存变更
    cacheChangeList   []CacheChangeInfo 

}

func NewTaobaoXhotelIntlAriNotifyRequest() *TaobaoXhotelIntlAriNotifyRequest{
    return &TaobaoXhotelIntlAriNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelIntlAriNotifyRequest) GetApiMethodName() string {
    return "taobao.xhotel.intl.ari.notify"
}

func (r TaobaoXhotelIntlAriNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelIntlAriNotifyRequest) SetCacheChangeList(cacheChangeList []CacheChangeInfo) error {
    r.cacheChangeList = cacheChangeList
    r.Set("cache_change_list", cacheChangeList)
    return nil
}

func (r TaobaoXhotelIntlAriNotifyRequest) GetCacheChangeList() []CacheChangeInfo {
    return r.cacheChangeList
}

