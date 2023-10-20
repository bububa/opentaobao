package tuanhotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelItemSkuUpdateAPIRequest 酒店团购套餐商品SKU更新和新增 API请求
// alitrip.tuan.hotel.item.sku.update
//
// 商户对发布的宝贝套餐价格库存信息进行更新，对于已存在的sku，未进行传递则不会进行覆盖。skuId必须为已存在的skuId，暂不支持库存类型的更改。因发布页改造升级，2020.03.05将下线此接口的新增SKU功能，更新SKU功能将保留，但商户2020.03.05后须前往发布页进行宝贝更新后，方可调用本接口。对于日历库存宝贝日历维度的价格和库存数据的更新，此接口存在调用超时的问题，不推荐使用，若有诉求，请使用alitrip.tuan.hotel.item.sku.calendar.update接口（该接口提供增量更新能力），接口地址为https://open.taobao.com/api.htm?docId=48160&amp;docType=2&amp;scopeId=12326
type AlitripTuanHotelItemSkuUpdateAPIRequest struct {
	model.Params
	// 关于sku（价格策略）的字段填写的说明  国内酒店套餐类目(日历库存必填选项：套餐名称、原价、间夜;普通库存必填选项：套餐名称、价格、原价、库存、间夜)。  国际酒店套餐类目(日历库存必填选型：套餐名称、原价、间夜、人数;普通库存必填选项：套餐名称、价格、原件、库存、间夜、人数)。  酒店餐饮美食类目(日历库存必填选项：套餐名称、原价、人数、次数;普通库存必填选项：套餐名称、价格、原价、库存、人数，次数)。  酒店服务类目(日历库存必填选项：套餐名称、原价、使用次数;普通库存必填选项：套餐名称、价格、原价、库存、使用次数)。  酒店客房优惠券类目(无sku（价格策略）选项，不填写)。
	_itemSkuList []TopTuanItemSkuVo
	// 宝贝ID
	_itemId int64
	// 宝贝所属类目
	_catId int64
}

// NewAlitripTuanHotelItemSkuUpdateRequest 初始化AlitripTuanHotelItemSkuUpdateAPIRequest对象
func NewAlitripTuanHotelItemSkuUpdateRequest() *AlitripTuanHotelItemSkuUpdateAPIRequest {
	return &AlitripTuanHotelItemSkuUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTuanHotelItemSkuUpdateAPIRequest) Reset() {
	r._itemSkuList = r._itemSkuList[:0]
	r._itemId = 0
	r._catId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelItemSkuUpdateAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.item.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelItemSkuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTuanHotelItemSkuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemSkuList is ItemSkuList Setter
// 关于sku（价格策略）的字段填写的说明  国内酒店套餐类目(日历库存必填选项：套餐名称、原价、间夜;普通库存必填选项：套餐名称、价格、原价、库存、间夜)。  国际酒店套餐类目(日历库存必填选型：套餐名称、原价、间夜、人数;普通库存必填选项：套餐名称、价格、原件、库存、间夜、人数)。  酒店餐饮美食类目(日历库存必填选项：套餐名称、原价、人数、次数;普通库存必填选项：套餐名称、价格、原价、库存、人数，次数)。  酒店服务类目(日历库存必填选项：套餐名称、原价、使用次数;普通库存必填选项：套餐名称、价格、原价、库存、使用次数)。  酒店客房优惠券类目(无sku（价格策略）选项，不填写)。
func (r *AlitripTuanHotelItemSkuUpdateAPIRequest) SetItemSkuList(_itemSkuList []TopTuanItemSkuVo) error {
	r._itemSkuList = _itemSkuList
	r.Set("item_sku_list", _itemSkuList)
	return nil
}

// GetItemSkuList ItemSkuList Getter
func (r AlitripTuanHotelItemSkuUpdateAPIRequest) GetItemSkuList() []TopTuanItemSkuVo {
	return r._itemSkuList
}

// SetItemId is ItemId Setter
// 宝贝ID
func (r *AlitripTuanHotelItemSkuUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripTuanHotelItemSkuUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCatId is CatId Setter
// 宝贝所属类目
func (r *AlitripTuanHotelItemSkuUpdateAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlitripTuanHotelItemSkuUpdateAPIRequest) GetCatId() int64 {
	return r._catId
}

var poolAlitripTuanHotelItemSkuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTuanHotelItemSkuUpdateRequest()
	},
}

// GetAlitripTuanHotelItemSkuUpdateRequest 从 sync.Pool 获取 AlitripTuanHotelItemSkuUpdateAPIRequest
func GetAlitripTuanHotelItemSkuUpdateAPIRequest() *AlitripTuanHotelItemSkuUpdateAPIRequest {
	return poolAlitripTuanHotelItemSkuUpdateAPIRequest.Get().(*AlitripTuanHotelItemSkuUpdateAPIRequest)
}

// ReleaseAlitripTuanHotelItemSkuUpdateAPIRequest 将 AlitripTuanHotelItemSkuUpdateAPIRequest 放入 sync.Pool
func ReleaseAlitripTuanHotelItemSkuUpdateAPIRequest(v *AlitripTuanHotelItemSkuUpdateAPIRequest) {
	v.Reset()
	poolAlitripTuanHotelItemSkuUpdateAPIRequest.Put(v)
}
