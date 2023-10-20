package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkprivilegegetAPIRequest 淘宝客-服务商-单品券高效转链 API请求
// taobao.tbk.privilege.get
//
// 单品券高效转链API
type TaobaotbkprivilegegetAPIRequest struct {
	model.Params
	// 渠道关系ID，仅适用于渠道推广场景（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）
	_relationid string
	// 会员运营ID
	_specialid string
	// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
	_externalid string
	// 团长与下游渠道合作的特殊标识，用于统计渠道推广效果
	_xid string
	// 商品ID
	_itemid string
	// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
	_promotiontype string
	// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）
	_bizsceneid string
	// 推广位id，mm_xx_xx_xx pid三段式中的第三段
	_adzoneid int64
	// 1：PC，2：无线，默认：１
	_platform int64
	// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
	_siteid int64
	// 会员人群ID，用于统计人群推广效果
	_ucrowdid int64
	// 是否获取前N件佣金 ,0-否，1-是,其他值-否
	_gettopnrate int64
	// 是否需要获取小程序链接，需要设置1。(暂未对外开放)
	_miniprogramlink int64
	// 商品库服务账户(场景id3权限对应的memberid）
	_manageitempubid int64
	// 入参商品id下的skuid，传入时会透传至转链结果url中
	_skuid int64
}

// NewTaobaotbkprivilegegetRequest 初始化TaobaotbkprivilegegetAPIRequest对象
func NewTaobaotbkprivilegegetRequest() *TaobaotbkprivilegegetAPIRequest {
	return &TaobaotbkprivilegegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkprivilegegetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.privilege.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkprivilegegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkprivilegegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationid is Relationid Setter
// 渠道关系ID，仅适用于渠道推广场景（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）
func (r *TaobaotbkprivilegegetAPIRequest) SetRelationid(_relationid string) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetRelationid() string {
	return r._relationid
}

// SetSpecialid is Specialid Setter
// 会员运营ID
func (r *TaobaotbkprivilegegetAPIRequest) SetSpecialid(_specialid string) error {
	r._specialid = _specialid
	r.Set("special_id", _specialid)
	return nil
}

// GetSpecialid Specialid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetSpecialid() string {
	return r._specialid
}

// SetExternalid is Externalid Setter
// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
func (r *TaobaotbkprivilegegetAPIRequest) SetExternalid(_externalid string) error {
	r._externalid = _externalid
	r.Set("external_id", _externalid)
	return nil
}

// GetExternalid Externalid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetExternalid() string {
	return r._externalid
}

// SetXid is Xid Setter
// 团长与下游渠道合作的特殊标识，用于统计渠道推广效果
func (r *TaobaotbkprivilegegetAPIRequest) SetXid(_xid string) error {
	r._xid = _xid
	r.Set("xid", _xid)
	return nil
}

// GetXid Xid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetXid() string {
	return r._xid
}

// SetItemid is Itemid Setter
// 商品ID
func (r *TaobaotbkprivilegegetAPIRequest) SetItemid(_itemid string) error {
	r._itemid = _itemid
	r.Set("item_id", _itemid)
	return nil
}

// GetItemid Itemid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetItemid() string {
	return r._itemid
}

// SetPromotiontype is Promotiontype Setter
// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
func (r *TaobaotbkprivilegegetAPIRequest) SetPromotiontype(_promotiontype string) error {
	r._promotiontype = _promotiontype
	r.Set("promotion_type", _promotiontype)
	return nil
}

// GetPromotiontype Promotiontype Getter
func (r TaobaotbkprivilegegetAPIRequest) GetPromotiontype() string {
	return r._promotiontype
}

// SetBizsceneid is Bizsceneid Setter
// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）
func (r *TaobaotbkprivilegegetAPIRequest) SetBizsceneid(_bizsceneid string) error {
	r._bizsceneid = _bizsceneid
	r.Set("biz_scene_id", _bizsceneid)
	return nil
}

// GetBizsceneid Bizsceneid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetBizsceneid() string {
	return r._bizsceneid
}

// SetAdzoneid is Adzoneid Setter
// 推广位id，mm_xx_xx_xx pid三段式中的第三段
func (r *TaobaotbkprivilegegetAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}

// SetPlatform is Platform Setter
// 1：PC，2：无线，默认：１
func (r *TaobaotbkprivilegegetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaotbkprivilegegetAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetSiteid is Siteid Setter
// 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
func (r *TaobaotbkprivilegegetAPIRequest) SetSiteid(_siteid int64) error {
	r._siteid = _siteid
	r.Set("site_id", _siteid)
	return nil
}

// GetSiteid Siteid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetSiteid() int64 {
	return r._siteid
}

// SetUcrowdid is Ucrowdid Setter
// 会员人群ID，用于统计人群推广效果
func (r *TaobaotbkprivilegegetAPIRequest) SetUcrowdid(_ucrowdid int64) error {
	r._ucrowdid = _ucrowdid
	r.Set("ucrowd_id", _ucrowdid)
	return nil
}

// GetUcrowdid Ucrowdid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetUcrowdid() int64 {
	return r._ucrowdid
}

// SetGettopnrate is Gettopnrate Setter
// 是否获取前N件佣金 ,0-否，1-是,其他值-否
func (r *TaobaotbkprivilegegetAPIRequest) SetGettopnrate(_gettopnrate int64) error {
	r._gettopnrate = _gettopnrate
	r.Set("get_topn_rate", _gettopnrate)
	return nil
}

// GetGettopnrate Gettopnrate Getter
func (r TaobaotbkprivilegegetAPIRequest) GetGettopnrate() int64 {
	return r._gettopnrate
}

// SetMiniprogramlink is Miniprogramlink Setter
// 是否需要获取小程序链接，需要设置1。(暂未对外开放)
func (r *TaobaotbkprivilegegetAPIRequest) SetMiniprogramlink(_miniprogramlink int64) error {
	r._miniprogramlink = _miniprogramlink
	r.Set("mini_program_link", _miniprogramlink)
	return nil
}

// GetMiniprogramlink Miniprogramlink Getter
func (r TaobaotbkprivilegegetAPIRequest) GetMiniprogramlink() int64 {
	return r._miniprogramlink
}

// SetManageitempubid is Manageitempubid Setter
// 商品库服务账户(场景id3权限对应的memberid）
func (r *TaobaotbkprivilegegetAPIRequest) SetManageitempubid(_manageitempubid int64) error {
	r._manageitempubid = _manageitempubid
	r.Set("manage_item_pub_id", _manageitempubid)
	return nil
}

// GetManageitempubid Manageitempubid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetManageitempubid() int64 {
	return r._manageitempubid
}

// SetSkuid is Skuid Setter
// 入参商品id下的skuid，传入时会透传至转链结果url中
func (r *TaobaotbkprivilegegetAPIRequest) SetSkuid(_skuid int64) error {
	r._skuid = _skuid
	r.Set("sku_id", _skuid)
	return nil
}

// GetSkuid Skuid Getter
func (r TaobaotbkprivilegegetAPIRequest) GetSkuid() int64 {
	return r._skuid
}
