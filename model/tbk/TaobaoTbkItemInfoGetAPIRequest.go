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
	_numIids string
	// ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	_ip string
	// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）
	_bizSceneId string
	// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
	_promotionType string
	// 渠道关系ID
	_relationId string
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 商品库服务账户(场景id3权限对应的memberid）
	_manageItemPubId int64
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

// SetNumIids is NumIids Setter
// 商品ID串，用,分割，最大40个
func (r *TaobaotbkiteminfogetAPIRequest) SetNumIids(_numIids string) error {
	r._numIids = _numIids
	r.Set("num_iids", _numIids)
	return nil
}

// GetNumIids NumIids Getter
func (r TaobaotbkiteminfogetAPIRequest) GetNumIids() string {
	return r._numIids
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

// SetBizSceneId is BizSceneId Setter
// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）
func (r *TaobaotbkiteminfogetAPIRequest) SetBizSceneId(_bizSceneId string) error {
	r._bizSceneId = _bizSceneId
	r.Set("biz_scene_id", _bizSceneId)
	return nil
}

// GetBizSceneId BizSceneId Getter
func (r TaobaotbkiteminfogetAPIRequest) GetBizSceneId() string {
	return r._bizSceneId
}

// SetPromotionType is PromotionType Setter
// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
func (r *TaobaotbkiteminfogetAPIRequest) SetPromotionType(_promotionType string) error {
	r._promotionType = _promotionType
	r.Set("promotion_type", _promotionType)
	return nil
}

// GetPromotionType PromotionType Getter
func (r TaobaotbkiteminfogetAPIRequest) GetPromotionType() string {
	return r._promotionType
}

// SetRelationId is RelationId Setter
// 渠道关系ID
func (r *TaobaotbkiteminfogetAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaotbkiteminfogetAPIRequest) GetRelationId() string {
	return r._relationId
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

// SetManageItemPubId is ManageItemPubId Setter
// 商品库服务账户(场景id3权限对应的memberid）
func (r *TaobaotbkiteminfogetAPIRequest) SetManageItemPubId(_manageItemPubId int64) error {
	r._manageItemPubId = _manageItemPubId
	r.Set("manage_item_pub_id", _manageItemPubId)
	return nil
}

// GetManageItemPubId ManageItemPubId Getter
func (r TaobaotbkiteminfogetAPIRequest) GetManageItemPubId() int64 {
	return r._manageItemPubId
}
