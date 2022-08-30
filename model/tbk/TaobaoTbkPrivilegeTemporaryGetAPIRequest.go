package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkPrivilegeTemporaryGetAPIRequest 淘宝客-服务商-单品券高效转链（临时接口） API请求
// taobao.tbk.privilege.temporary.get
//
// 单品券高效转链API
type TaobaoTbkPrivilegeTemporaryGetAPIRequest struct {
	model.Params
	// 渠道关系ID，仅适用于渠道推广场景
	_relationId string
	// 会员运营ID
	_specialId string
	// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
	_externalId string
	// 团长与下游渠道合作的特殊标识，用于统计渠道推广效果
	_xid string
	// 淘客商品id
	_itemId string
	// 推广位id，mm_xx_xx_xx pid三段式中的第三段
	_adzoneId int64
	// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
	_siteId int64
	// 1：PC，2：无线，默认：１
	_platform int64
	// 会员人群ID，用于统计人群推广效果
	_ucrowdId int64
	// 是否获取前N件佣金 ,0-否，1-是,其他值-否
	_getTopnRate int64
	// 是否需要获取小程序链接，需要设置1。(暂未对外开放)
	_miniProgramLink int64
}

// NewTaobaoTbkPrivilegeTemporaryGetRequest 初始化TaobaoTbkPrivilegeTemporaryGetAPIRequest对象
func NewTaobaoTbkPrivilegeTemporaryGetRequest() *TaobaoTbkPrivilegeTemporaryGetAPIRequest {
	return &TaobaoTbkPrivilegeTemporaryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.privilege.temporary.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRelationId is RelationId Setter
// 渠道关系ID，仅适用于渠道推广场景
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetSpecialId is SpecialId Setter
// 会员运营ID
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetExternalId is ExternalId Setter
// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetExternalId(_externalId string) error {
	r._externalId = _externalId
	r.Set("external_id", _externalId)
	return nil
}

// GetExternalId ExternalId Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetExternalId() string {
	return r._externalId
}

// SetXid is Xid Setter
// 团长与下游渠道合作的特殊标识，用于统计渠道推广效果
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetXid(_xid string) error {
	r._xid = _xid
	r.Set("xid", _xid)
	return nil
}

// GetXid Xid Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetXid() string {
	return r._xid
}

// SetItemId is ItemId Setter
// 淘客商品id
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetItemId() string {
	return r._itemId
}

// SetAdzoneId is AdzoneId Setter
// 推广位id，mm_xx_xx_xx pid三段式中的第三段
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetSiteId is SiteId Setter
// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetSiteId() int64 {
	return r._siteId
}

// SetPlatform is Platform Setter
// 1：PC，2：无线，默认：１
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetUcrowdId is UcrowdId Setter
// 会员人群ID，用于统计人群推广效果
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}

// SetGetTopnRate is GetTopnRate Setter
// 是否获取前N件佣金 ,0-否，1-是,其他值-否
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetGetTopnRate(_getTopnRate int64) error {
	r._getTopnRate = _getTopnRate
	r.Set("get_topn_rate", _getTopnRate)
	return nil
}

// GetGetTopnRate GetTopnRate Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetGetTopnRate() int64 {
	return r._getTopnRate
}

// SetMiniProgramLink is MiniProgramLink Setter
// 是否需要获取小程序链接，需要设置1。(暂未对外开放)
func (r *TaobaoTbkPrivilegeTemporaryGetAPIRequest) SetMiniProgramLink(_miniProgramLink int64) error {
	r._miniProgramLink = _miniProgramLink
	r.Set("mini_program_link", _miniProgramLink)
	return nil
}

// GetMiniProgramLink MiniProgramLink Getter
func (r TaobaoTbkPrivilegeTemporaryGetAPIRequest) GetMiniProgramLink() int64 {
	return r._miniProgramLink
}
