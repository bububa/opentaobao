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
    platform   string
    // 达人ID
    adUserId   int64
    // 站外是1
    sourceType   int64
    // 业务方，新浪为sina
    bizType   string
    // 商品ID
    itemId   int64
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
func (r *TaobaoBeehiveItemCpsUrlRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetPlatform() string {
    return r.platform
}
// AdUserId Setter
// 达人ID
func (r *TaobaoBeehiveItemCpsUrlRequest) SetAdUserId(adUserId int64) error {
    r.adUserId = adUserId
    r.Set("ad_user_id", adUserId)
    return nil
}

// AdUserId Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetAdUserId() int64 {
    return r.adUserId
}
// SourceType Setter
// 站外是1
func (r *TaobaoBeehiveItemCpsUrlRequest) SetSourceType(sourceType int64) error {
    r.sourceType = sourceType
    r.Set("source_type", sourceType)
    return nil
}

// SourceType Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetSourceType() int64 {
    return r.sourceType
}
// BizType Setter
// 业务方，新浪为sina
func (r *TaobaoBeehiveItemCpsUrlRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetBizType() string {
    return r.bizType
}
// ItemId Setter
// 商品ID
func (r *TaobaoBeehiveItemCpsUrlRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoBeehiveItemCpsUrlRequest) GetItemId() int64 {
    return r.itemId
}
