package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelIntlAriNotifyAPIRequest 国际酒店集团价库变更通知 API请求
// taobao.xhotel.intl.ari.notify
//
// 国际酒店集团价库变更时通知变更内容，平台及时更新价库信息，保证价库新鲜度
type TaobaoXhotelIntlAriNotifyAPIRequest struct {
	model.Params
	// 缓存变更
	_cacheChangeList []CacheChangeInfo
}

// NewTaobaoXhotelIntlAriNotifyRequest 初始化TaobaoXhotelIntlAriNotifyAPIRequest对象
func NewTaobaoXhotelIntlAriNotifyRequest() *TaobaoXhotelIntlAriNotifyAPIRequest {
	return &TaobaoXhotelIntlAriNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelIntlAriNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.intl.ari.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelIntlAriNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCacheChangeList is CacheChangeList Setter
// 缓存变更
func (r *TaobaoXhotelIntlAriNotifyAPIRequest) SetCacheChangeList(_cacheChangeList []CacheChangeInfo) error {
	r._cacheChangeList = _cacheChangeList
	r.Set("cache_change_list", _cacheChangeList)
	return nil
}

// GetCacheChangeList CacheChangeList Getter
func (r TaobaoXhotelIntlAriNotifyAPIRequest) GetCacheChangeList() []CacheChangeInfo {
	return r._cacheChangeList
}
