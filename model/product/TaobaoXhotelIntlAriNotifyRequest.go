package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际酒店集团价库变更通知 API请求
taobao.xhotel.intl.ari.notify

国际酒店集团价库变更时通知变更内容，平台及时更新价库信息，保证价库新鲜度
*/
type TaobaoXhotelIntlAriNotifyRequest struct {
    model.Params
    // 缓存变更
    cacheChangeList   []CacheChangeInfo
}

// 初始化TaobaoXhotelIntlAriNotifyRequest对象
func NewTaobaoXhotelIntlAriNotifyRequest() *TaobaoXhotelIntlAriNotifyRequest{
    return &TaobaoXhotelIntlAriNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelIntlAriNotifyRequest) GetApiMethodName() string {
    return "taobao.xhotel.intl.ari.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelIntlAriNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CacheChangeList Setter
// 缓存变更
func (r *TaobaoXhotelIntlAriNotifyRequest) SetCacheChangeList(cacheChangeList []CacheChangeInfo) error {
    r.cacheChangeList = cacheChangeList
    r.Set("cache_change_list", cacheChangeList)
    return nil
}

// CacheChangeList Getter
func (r TaobaoXhotelIntlAriNotifyRequest) GetCacheChangeList() []CacheChangeInfo {
    return r.cacheChangeList
}
