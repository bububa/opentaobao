package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScGeneralLinkConvertAPIRequest 淘宝客-服务商-万能转链 API请求
// taobao.tbk.sc.general.link.convert
//
// 淘宝客-服务商-万能转链
type TaobaoTbkScGeneralLinkConvertAPIRequest struct {
	model.Params
	// 商品转链
	_itemDto []LkItemDto
	// 会场页面转链
	_pageDto []LkPageDto
	// 店铺转链
	_shopDto []LkShopDto
	// 链接/口令转链
	_materialDto []LkMaterialDto
	// 1-动态ID转链场景，2-消费者比价场景，4-
	_bizSceneId string
	// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，其它模式不用填写本字段）
	_promotionType string
	// 物料列表，可以为url或淘口令,多个时使用英文逗号拼接传入
	_materialList string
	// 渠道管理ID（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）
	_relationId string
	// 卖家ID列表,多个时使用英文逗号拼接传入
	_sellerIdList string
	// 商品ID列表,多个时使用英文逗号拼接传入
	_itemIdList string
	// 会场ID列表,多个时使用英文逗号拼接传入
	_pageIdList string
	// 会员运营id
	_specialId string
	// 加密用户标识
	_uvid string
	// 启明系统任务ID
	_qmtid string
	// 推广位id，mm_xx_xx_xx pid三段式中的第三段
	_adzoneId int64
	// 推广位id，mm_xx_xx_xx pid三段式中的第二段
	_siteId int64
	// 会场页面内定坑商品
	_targetItem *TargetItemDto
}

// NewTaobaoTbkScGeneralLinkConvertRequest 初始化TaobaoTbkScGeneralLinkConvertAPIRequest对象
func NewTaobaoTbkScGeneralLinkConvertRequest() *TaobaoTbkScGeneralLinkConvertAPIRequest {
	return &TaobaoTbkScGeneralLinkConvertAPIRequest{
		Params: model.NewParams(17),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) Reset() {
	r._itemDto = r._itemDto[:0]
	r._pageDto = r._pageDto[:0]
	r._shopDto = r._shopDto[:0]
	r._materialDto = r._materialDto[:0]
	r._bizSceneId = ""
	r._promotionType = ""
	r._materialList = ""
	r._relationId = ""
	r._sellerIdList = ""
	r._itemIdList = ""
	r._pageIdList = ""
	r._specialId = ""
	r._uvid = ""
	r._qmtid = ""
	r._adzoneId = 0
	r._siteId = 0
	r._targetItem = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.general.link.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDto is ItemDto Setter
// 商品转链
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetItemDto(_itemDto []LkItemDto) error {
	r._itemDto = _itemDto
	r.Set("item_dto", _itemDto)
	return nil
}

// GetItemDto ItemDto Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetItemDto() []LkItemDto {
	return r._itemDto
}

// SetPageDto is PageDto Setter
// 会场页面转链
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetPageDto(_pageDto []LkPageDto) error {
	r._pageDto = _pageDto
	r.Set("page_dto", _pageDto)
	return nil
}

// GetPageDto PageDto Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetPageDto() []LkPageDto {
	return r._pageDto
}

// SetShopDto is ShopDto Setter
// 店铺转链
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetShopDto(_shopDto []LkShopDto) error {
	r._shopDto = _shopDto
	r.Set("shop_dto", _shopDto)
	return nil
}

// GetShopDto ShopDto Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetShopDto() []LkShopDto {
	return r._shopDto
}

// SetMaterialDto is MaterialDto Setter
// 链接/口令转链
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetMaterialDto(_materialDto []LkMaterialDto) error {
	r._materialDto = _materialDto
	r.Set("material_dto", _materialDto)
	return nil
}

// GetMaterialDto MaterialDto Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetMaterialDto() []LkMaterialDto {
	return r._materialDto
}

// SetBizSceneId is BizSceneId Setter
// 1-动态ID转链场景，2-消费者比价场景，4-
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetBizSceneId(_bizSceneId string) error {
	r._bizSceneId = _bizSceneId
	r.Set("biz_scene_id", _bizSceneId)
	return nil
}

// GetBizSceneId BizSceneId Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetBizSceneId() string {
	return r._bizSceneId
}

// SetPromotionType is PromotionType Setter
// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，其它模式不用填写本字段）
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetPromotionType(_promotionType string) error {
	r._promotionType = _promotionType
	r.Set("promotion_type", _promotionType)
	return nil
}

// GetPromotionType PromotionType Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetPromotionType() string {
	return r._promotionType
}

// SetMaterialList is MaterialList Setter
// 物料列表，可以为url或淘口令,多个时使用英文逗号拼接传入
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetMaterialList(_materialList string) error {
	r._materialList = _materialList
	r.Set("material_list", _materialList)
	return nil
}

// GetMaterialList MaterialList Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetMaterialList() string {
	return r._materialList
}

// SetRelationId is RelationId Setter
// 渠道管理ID（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetSellerIdList is SellerIdList Setter
// 卖家ID列表,多个时使用英文逗号拼接传入
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetSellerIdList(_sellerIdList string) error {
	r._sellerIdList = _sellerIdList
	r.Set("seller_id_list", _sellerIdList)
	return nil
}

// GetSellerIdList SellerIdList Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetSellerIdList() string {
	return r._sellerIdList
}

// SetItemIdList is ItemIdList Setter
// 商品ID列表,多个时使用英文逗号拼接传入
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetItemIdList(_itemIdList string) error {
	r._itemIdList = _itemIdList
	r.Set("item_id_list", _itemIdList)
	return nil
}

// GetItemIdList ItemIdList Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetItemIdList() string {
	return r._itemIdList
}

// SetPageIdList is PageIdList Setter
// 会场ID列表,多个时使用英文逗号拼接传入
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetPageIdList(_pageIdList string) error {
	r._pageIdList = _pageIdList
	r.Set("page_id_list", _pageIdList)
	return nil
}

// GetPageIdList PageIdList Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetPageIdList() string {
	return r._pageIdList
}

// SetSpecialId is SpecialId Setter
// 会员运营id
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetUvid is Uvid Setter
// 加密用户标识
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetUvid(_uvid string) error {
	r._uvid = _uvid
	r.Set("uvid", _uvid)
	return nil
}

// GetUvid Uvid Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetUvid() string {
	return r._uvid
}

// SetQmtid is Qmtid Setter
// 启明系统任务ID
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetQmtid(_qmtid string) error {
	r._qmtid = _qmtid
	r.Set("qmtid", _qmtid)
	return nil
}

// GetQmtid Qmtid Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetQmtid() string {
	return r._qmtid
}

// SetAdzoneId is AdzoneId Setter
// 推广位id，mm_xx_xx_xx pid三段式中的第三段
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetSiteId is SiteId Setter
// 推广位id，mm_xx_xx_xx pid三段式中的第二段
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetSiteId() int64 {
	return r._siteId
}

// SetTargetItem is TargetItem Setter
// 会场页面内定坑商品
func (r *TaobaoTbkScGeneralLinkConvertAPIRequest) SetTargetItem(_targetItem *TargetItemDto) error {
	r._targetItem = _targetItem
	r.Set("target_item", _targetItem)
	return nil
}

// GetTargetItem TargetItem Getter
func (r TaobaoTbkScGeneralLinkConvertAPIRequest) GetTargetItem() *TargetItemDto {
	return r._targetItem
}

var poolTaobaoTbkScGeneralLinkConvertAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScGeneralLinkConvertRequest()
	},
}

// GetTaobaoTbkScGeneralLinkConvertRequest 从 sync.Pool 获取 TaobaoTbkScGeneralLinkConvertAPIRequest
func GetTaobaoTbkScGeneralLinkConvertAPIRequest() *TaobaoTbkScGeneralLinkConvertAPIRequest {
	return poolTaobaoTbkScGeneralLinkConvertAPIRequest.Get().(*TaobaoTbkScGeneralLinkConvertAPIRequest)
}

// ReleaseTaobaoTbkScGeneralLinkConvertAPIRequest 将 TaobaoTbkScGeneralLinkConvertAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScGeneralLinkConvertAPIRequest(v *TaobaoTbkScGeneralLinkConvertAPIRequest) {
	v.Reset()
	poolTaobaoTbkScGeneralLinkConvertAPIRequest.Put(v)
}
