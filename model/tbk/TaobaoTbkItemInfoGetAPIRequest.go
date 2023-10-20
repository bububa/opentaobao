package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkItemInfoGetAPIRequest 淘宝客-公用-淘宝客商品详情查询(简版) API请求
// taobao.tbk.item.info.get
//
// 淘宝客商品详情查询（简版）
type TaobaoTbkItemInfoGetAPIRequest struct {
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

// NewTaobaoTbkItemInfoGetRequest 初始化TaobaoTbkItemInfoGetAPIRequest对象
func NewTaobaoTbkItemInfoGetRequest() *TaobaoTbkItemInfoGetAPIRequest {
	return &TaobaoTbkItemInfoGetAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkItemInfoGetAPIRequest) Reset() {
	r._numIids = ""
	r._ip = ""
	r._bizSceneId = ""
	r._promotionType = ""
	r._relationId = ""
	r._platform = 0
	r._manageItemPubId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkItemInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.item.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkItemInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkItemInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIids is NumIids Setter
// 商品ID串，用,分割，最大40个
func (r *TaobaoTbkItemInfoGetAPIRequest) SetNumIids(_numIids string) error {
	r._numIids = _numIids
	r.Set("num_iids", _numIids)
	return nil
}

// GetNumIids NumIids Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetNumIids() string {
	return r._numIids
}

// SetIp is Ip Setter
// ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
func (r *TaobaoTbkItemInfoGetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetIp() string {
	return r._ip
}

// SetBizSceneId is BizSceneId Setter
// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景（不填默认为1）
func (r *TaobaoTbkItemInfoGetAPIRequest) SetBizSceneId(_bizSceneId string) error {
	r._bizSceneId = _bizSceneId
	r.Set("biz_scene_id", _bizSceneId)
	return nil
}

// GetBizSceneId BizSceneId Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetBizSceneId() string {
	return r._bizSceneId
}

// SetPromotionType is PromotionType Setter
// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
func (r *TaobaoTbkItemInfoGetAPIRequest) SetPromotionType(_promotionType string) error {
	r._promotionType = _promotionType
	r.Set("promotion_type", _promotionType)
	return nil
}

// GetPromotionType PromotionType Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetPromotionType() string {
	return r._promotionType
}

// SetRelationId is RelationId Setter
// 渠道关系ID
func (r *TaobaoTbkItemInfoGetAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaoTbkItemInfoGetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetManageItemPubId is ManageItemPubId Setter
// 商品库服务账户(场景id3权限对应的memberid）
func (r *TaobaoTbkItemInfoGetAPIRequest) SetManageItemPubId(_manageItemPubId int64) error {
	r._manageItemPubId = _manageItemPubId
	r.Set("manage_item_pub_id", _manageItemPubId)
	return nil
}

// GetManageItemPubId ManageItemPubId Getter
func (r TaobaoTbkItemInfoGetAPIRequest) GetManageItemPubId() int64 {
	return r._manageItemPubId
}

var poolTaobaoTbkItemInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkItemInfoGetRequest()
	},
}

// GetTaobaoTbkItemInfoGetRequest 从 sync.Pool 获取 TaobaoTbkItemInfoGetAPIRequest
func GetTaobaoTbkItemInfoGetAPIRequest() *TaobaoTbkItemInfoGetAPIRequest {
	return poolTaobaoTbkItemInfoGetAPIRequest.Get().(*TaobaoTbkItemInfoGetAPIRequest)
}

// ReleaseTaobaoTbkItemInfoGetAPIRequest 将 TaobaoTbkItemInfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkItemInfoGetAPIRequest(v *TaobaoTbkItemInfoGetAPIRequest) {
	v.Reset()
	poolTaobaoTbkItemInfoGetAPIRequest.Put(v)
}
