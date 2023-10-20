package beehive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobeehiveitemcpsurlAPIRequest 分佣链接生成接口 API请求
// taobao.beehive.item.cps.url
//
// 传入包括itemId,accountId,bizType在内的参数，对应参数返回分佣链接
type TaobaobeehiveitemcpsurlAPIRequest struct {
	model.Params
	// 平台，一般为手机
	_platform string
	// 业务方，新浪为sina
	_bizType string
	// 商品ID
	_itemId int64
	// 达人ID
	_adUserId int64
	// 站外是1
	_sourceType int64
}

// NewTaobaobeehiveitemcpsurlRequest 初始化TaobaobeehiveitemcpsurlAPIRequest对象
func NewTaobaobeehiveitemcpsurlRequest() *TaobaobeehiveitemcpsurlAPIRequest {
	return &TaobaobeehiveitemcpsurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobeehiveitemcpsurlAPIRequest) GetApiMethodName() string {
	return "taobao.beehive.item.cps.url"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobeehiveitemcpsurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobeehiveitemcpsurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlatform is Platform Setter
// 平台，一般为手机
func (r *TaobaobeehiveitemcpsurlAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaobeehiveitemcpsurlAPIRequest) GetPlatform() string {
	return r._platform
}

// SetBizType is BizType Setter
// 业务方，新浪为sina
func (r *TaobaobeehiveitemcpsurlAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaobeehiveitemcpsurlAPIRequest) GetBizType() string {
	return r._bizType
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaobeehiveitemcpsurlAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaobeehiveitemcpsurlAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetAdUserId is AdUserId Setter
// 达人ID
func (r *TaobaobeehiveitemcpsurlAPIRequest) SetAdUserId(_adUserId int64) error {
	r._adUserId = _adUserId
	r.Set("ad_user_id", _adUserId)
	return nil
}

// GetAdUserId AdUserId Getter
func (r TaobaobeehiveitemcpsurlAPIRequest) GetAdUserId() int64 {
	return r._adUserId
}

// SetSourceType is SourceType Setter
// 站外是1
func (r *TaobaobeehiveitemcpsurlAPIRequest) SetSourceType(_sourceType int64) error {
	r._sourceType = _sourceType
	r.Set("source_type", _sourceType)
	return nil
}

// GetSourceType SourceType Getter
func (r TaobaobeehiveitemcpsurlAPIRequest) GetSourceType() int64 {
	return r._sourceType
}
