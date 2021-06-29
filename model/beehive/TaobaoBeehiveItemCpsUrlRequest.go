package beehive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分佣链接生成接口 API请求
taobao.beehive.item.cps.url

传入包括itemId,accountId,bizType在内的参数，对应参数返回分佣链接
*/
type TaobaoBeehiveItemCpsUrlRequest struct {
    model.Params
    // 平台，一般为手机
    _platform   string
    // 达人ID
    _adUserId   int64
    // 站外是1
    _sourceType   int64
    // 业务方，新浪为sina
    _bizType   string
    // 商品ID
    _itemId   int64
}

// 初始化TaobaoBeehiveItemCpsUrlRequest对象
func NewTaobaoBeehiveItemCpsUrlRequest() *TaobaoBeehiveItemCpsUrlRequest{
    return &TaobaoBeehiveItemCpsUrlRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBeehiveItemCpsUrlRequest) GetApiMethodName() string {
    return "taobao.beehive.item.cps.url"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBeehiveItemCpsUrlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Platform Setter
// 平台，一般为手机
func (r *TaobaoBeehiveItemCpsUrlRequest) SetPlatform(_platform string) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetPlatform() string {
    return r._platform
}
// AdUserId Setter
// 达人ID
func (r *TaobaoBeehiveItemCpsUrlRequest) SetAdUserId(_adUserId int64) error {
    r._adUserId = _adUserId
    r.Set("ad_user_id", _adUserId)
    return nil
}

// AdUserId Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetAdUserId() int64 {
    return r._adUserId
}
// SourceType Setter
// 站外是1
func (r *TaobaoBeehiveItemCpsUrlRequest) SetSourceType(_sourceType int64) error {
    r._sourceType = _sourceType
    r.Set("source_type", _sourceType)
    return nil
}

// SourceType Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetSourceType() int64 {
    return r._sourceType
}
// BizType Setter
// 业务方，新浪为sina
func (r *TaobaoBeehiveItemCpsUrlRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetBizType() string {
    return r._bizType
}
// ItemId Setter
// 商品ID
func (r *TaobaoBeehiveItemCpsUrlRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetItemId() int64 {
    return r._itemId
}
