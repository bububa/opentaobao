package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ECRM创建淘短链服务 API请求
taobao.crm.service.channel.shortlink.create

可生成店铺宝贝、店铺首页、活动链接、订单链接等4种可呼起手机淘宝APP至对应页面的淘短链。
*/
type TaobaoCrmServiceChannelShortlinkCreateRequest struct {
    model.Params
    // 淘短链名称（最多只能16个中文字符，类型为订单链接时传入订单ID）。
    shortLinkName   string
    // 淘短链类型：LT_ITEM（商品淘短链）,LT_SHOP（店铺首页淘短链）,LT_ACTIVITY（活动页淘短链）,LT_TRADE（订单详情页淘短链）。
    linkType   string
    // 类型为LT_ITEM时必须传入商品ID，类型为LT_SHOP时必须传入空值，类型为LT_ACTIVITY时传入活动页URL（URL必须是taobao.com,tmall.com,jaeapp.com这三个域名下的URL），类型为LT_TRADE时传入订单ID。
    shortLinkData   string
}

// 初始化TaobaoCrmServiceChannelShortlinkCreateRequest对象
func NewTaobaoCrmServiceChannelShortlinkCreateRequest() *TaobaoCrmServiceChannelShortlinkCreateRequest{
    return &TaobaoCrmServiceChannelShortlinkCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmServiceChannelShortlinkCreateRequest) GetApiMethodName() string {
    return "taobao.crm.service.channel.shortlink.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmServiceChannelShortlinkCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShortLinkName Setter
// 淘短链名称（最多只能16个中文字符，类型为订单链接时传入订单ID）。
func (r *TaobaoCrmServiceChannelShortlinkCreateRequest) SetShortLinkName(shortLinkName string) error {
    r.shortLinkName = shortLinkName
    r.Set("short_link_name", shortLinkName)
    return nil
}

// ShortLinkName Getter
func (r TaobaoCrmServiceChannelShortlinkCreateRequest) GetShortLinkName() string {
    return r.shortLinkName
}
// LinkType Setter
// 淘短链类型：LT_ITEM（商品淘短链）,LT_SHOP（店铺首页淘短链）,LT_ACTIVITY（活动页淘短链）,LT_TRADE（订单详情页淘短链）。
func (r *TaobaoCrmServiceChannelShortlinkCreateRequest) SetLinkType(linkType string) error {
    r.linkType = linkType
    r.Set("link_type", linkType)
    return nil
}

// LinkType Getter
func (r TaobaoCrmServiceChannelShortlinkCreateRequest) GetLinkType() string {
    return r.linkType
}
// ShortLinkData Setter
// 类型为LT_ITEM时必须传入商品ID，类型为LT_SHOP时必须传入空值，类型为LT_ACTIVITY时传入活动页URL（URL必须是taobao.com,tmall.com,jaeapp.com这三个域名下的URL），类型为LT_TRADE时传入订单ID。
func (r *TaobaoCrmServiceChannelShortlinkCreateRequest) SetShortLinkData(shortLinkData string) error {
    r.shortLinkData = shortLinkData
    r.Set("short_link_data", shortLinkData)
    return nil
}

// ShortLinkData Getter
func (r TaobaoCrmServiceChannelShortlinkCreateRequest) GetShortLinkData() string {
    return r.shortLinkData
}
