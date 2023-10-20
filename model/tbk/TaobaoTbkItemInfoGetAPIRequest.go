package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkiteminfogetAPIRequest 淘宝客-公用-淘宝客商品详情查询(简版) API请求
// taobao.tbk.item.info.get
//
// 淘宝客商品详情查询（简版）
type TaobaotbkiteminfogetAPIRequest struct {
	model.Params
	// 商品ID串，用,分割，最大40个
	_numiids string
	// ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	_ip string
	// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）
	_bizsceneid string
	// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
	_promotiontype string
	// 渠道关系ID
	_relationid string
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 商品库服务账户(场景id3权限对应的memberid）
	_manageitempubid int64
}

// NewTaobaotbkiteminfogetRequest 初始化TaobaotbkiteminfogetAPIRequest对象
func NewTaobaotbkiteminfogetRequest() *TaobaotbkiteminfogetAPIRequest {
	return &TaobaotbkiteminfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkiteminfogetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.item.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkiteminfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkiteminfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumiids is Numiids Setter
// 商品ID串，用,分割，最大40个
func (r *TaobaotbkiteminfogetAPIRequest) SetNumiids(_numiids string) error {
	r._numiids = _numiids
	r.Set("num_iids", _numiids)
	return nil
}

// GetNumiids Numiids Getter
func (r TaobaotbkiteminfogetAPIRequest) GetNumiids() string {
	return r._numiids
}

// SetIp is Ip Setter
// ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
func (r *TaobaotbkiteminfogetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r TaobaotbkiteminfogetAPIRequest) GetIp() string {
	return r._ip
}

// SetBizsceneid is Bizsceneid Setter
// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）
func (r *TaobaotbkiteminfogetAPIRequest) SetBizsceneid(_bizsceneid string) error {
	r._bizsceneid = _bizsceneid
	r.Set("biz_scene_id", _bizsceneid)
	return nil
}

// GetBizsceneid Bizsceneid Getter
func (r TaobaotbkiteminfogetAPIRequest) GetBizsceneid() string {
	return r._bizsceneid
}

// SetPromotiontype is Promotiontype Setter
// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
func (r *TaobaotbkiteminfogetAPIRequest) SetPromotiontype(_promotiontype string) error {
	r._promotiontype = _promotiontype
	r.Set("promotion_type", _promotiontype)
	return nil
}

// GetPromotiontype Promotiontype Getter
func (r TaobaotbkiteminfogetAPIRequest) GetPromotiontype() string {
	return r._promotiontype
}

// SetRelationid is Relationid Setter
// 渠道关系ID
func (r *TaobaotbkiteminfogetAPIRequest) SetRelationid(_relationid string) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkiteminfogetAPIRequest) GetRelationid() string {
	return r._relationid
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaotbkiteminfogetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaotbkiteminfogetAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetManageitempubid is Manageitempubid Setter
// 商品库服务账户(场景id3权限对应的memberid）
func (r *TaobaotbkiteminfogetAPIRequest) SetManageitempubid(_manageitempubid int64) error {
	r._manageitempubid = _manageitempubid
	r.Set("manage_item_pub_id", _manageitempubid)
	return nil
}

// GetManageitempubid Manageitempubid Getter
func (r TaobaotbkiteminfogetAPIRequest) GetManageitempubid() int64 {
	return r._manageitempubid
}
