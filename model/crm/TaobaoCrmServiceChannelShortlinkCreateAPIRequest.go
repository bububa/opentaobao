package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmservicechannelshortlinkcreateAPIRequest ECRM创建淘短链服务 API请求
// taobao.crm.service.channel.shortlink.create
//
// 可生成店铺宝贝、店铺首页、活动链接、订单链接等4种可呼起手机淘宝APP至对应页面的淘短链。
type TaobaocrmservicechannelshortlinkcreateAPIRequest struct {
	model.Params
	// 淘短链类型：LT_ITEM（商品淘短链）,LT_SHOP（店铺首页淘短链）,LT_ACTIVITY（活动页淘短链）,LT_TRADE（订单详情页淘短链）。
	_linkType string
	// 类型为LT_ITEM时必须传入商品ID，类型为LT_SHOP时必须传入空值，类型为LT_ACTIVITY时传入活动页URL（URL必须是taobao.com,tmall.com,jaeapp.com这三个域名下的URL），类型为LT_TRADE时传入订单ID。
	_shortLinkData string
	// 淘短链名称（最多只能16个中文字符，类型为订单链接时传入订单ID）。
	_shortLinkName string
}

// NewTaobaocrmservicechannelshortlinkcreateRequest 初始化TaobaocrmservicechannelshortlinkcreateAPIRequest对象
func NewTaobaocrmservicechannelshortlinkcreateRequest() *TaobaocrmservicechannelshortlinkcreateAPIRequest {
	return &TaobaocrmservicechannelshortlinkcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmservicechannelshortlinkcreateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.service.channel.shortlink.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmservicechannelshortlinkcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmservicechannelshortlinkcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLinkType is LinkType Setter
// 淘短链类型：LT_ITEM（商品淘短链）,LT_SHOP（店铺首页淘短链）,LT_ACTIVITY（活动页淘短链）,LT_TRADE（订单详情页淘短链）。
func (r *TaobaocrmservicechannelshortlinkcreateAPIRequest) SetLinkType(_linkType string) error {
	r._linkType = _linkType
	r.Set("link_type", _linkType)
	return nil
}

// GetLinkType LinkType Getter
func (r TaobaocrmservicechannelshortlinkcreateAPIRequest) GetLinkType() string {
	return r._linkType
}

// SetShortLinkData is ShortLinkData Setter
// 类型为LT_ITEM时必须传入商品ID，类型为LT_SHOP时必须传入空值，类型为LT_ACTIVITY时传入活动页URL（URL必须是taobao.com,tmall.com,jaeapp.com这三个域名下的URL），类型为LT_TRADE时传入订单ID。
func (r *TaobaocrmservicechannelshortlinkcreateAPIRequest) SetShortLinkData(_shortLinkData string) error {
	r._shortLinkData = _shortLinkData
	r.Set("short_link_data", _shortLinkData)
	return nil
}

// GetShortLinkData ShortLinkData Getter
func (r TaobaocrmservicechannelshortlinkcreateAPIRequest) GetShortLinkData() string {
	return r._shortLinkData
}

// SetShortLinkName is ShortLinkName Setter
// 淘短链名称（最多只能16个中文字符，类型为订单链接时传入订单ID）。
func (r *TaobaocrmservicechannelshortlinkcreateAPIRequest) SetShortLinkName(_shortLinkName string) error {
	r._shortLinkName = _shortLinkName
	r.Set("short_link_name", _shortLinkName)
	return nil
}

// GetShortLinkName ShortLinkName Getter
func (r TaobaocrmservicechannelshortlinkcreateAPIRequest) GetShortLinkName() string {
	return r._shortLinkName
}
