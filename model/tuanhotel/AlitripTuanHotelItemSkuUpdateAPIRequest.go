package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTuanHotelItemSkuUpdateAPIRequest
酒店团购套餐商品SKU更新和新增 API请求
alitrip.tuan.hotel.item.sku.update

商户对发布的宝贝套餐价格库存信息进行更新，对于已存在的sku，未进行传递则不会进行覆盖。skuId必须为已存在的skuId，暂不支持库存类型的更改。因发布页改造升级，2020.03.05将下线此接口的新增SKU功能，更新SKU功能将保留，但商户2020.03.05后须前往发布页进行宝贝更新后，方可调用本接口。对于日历库存宝贝日历维度的价格和库存数据的更新，此接口存在调用超时的问题，不推荐使用，若有诉求，请使用alitrip.tuan.hotel.item.sku.calendar.update接口（该接口提供增量更新能力），接口地址为https://open.taobao.com/api.htm?docId=48160&docType=2&scopeId=12326 */
type AlitripTuanHotelItemSkuUpdateAPIRequest struct {
	model.Params
	// 宝贝ID
	_itemId int64
	// 宝贝所属类目
	_catId int64
	// 关于sku（价格策略）的字段填写的说明  国内酒店套餐类目(日历库存必填选项：套餐名称、原价、间夜;普通库存必填选项：套餐名称、价格、原价、库存、间夜)。  国际酒店套餐类目(日历库存必填选型：套餐名称、原价、间夜、人数;普通库存必填选项：套餐名称、价格、原件、库存、间夜、人数)。  酒店餐饮美食类目(日历库存必填选项：套餐名称、原价、人数、次数;普通库存必填选项：套餐名称、价格、原价、库存、人数，次数)。  酒店服务类目(日历库存必填选项：套餐名称、原价、使用次数;普通库存必填选项：套餐名称、价格、原价、库存、使用次数)。  酒店客房优惠券类目(无sku（价格策略）选项，不填写)。
	_itemSkuList []TopTuanItemSkuVO
}

// New
