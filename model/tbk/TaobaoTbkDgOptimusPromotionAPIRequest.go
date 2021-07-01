package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgOptimusPromotionAPIRequest
淘宝客-推广者-权益物料精选 API请求
taobao.tbk.dg.optimus.promotion

推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。 */
type TaobaoTbkDgOptimusPromotionAPIRequest struct {
	model.Params
	// 页大小，一次请求请限制在10以内
	_pageSize int64
	// 第几页，默认：1
	_pageNum int64
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId int64
	// 官方提供的权益物料Id。有价券-37104、大额店铺券-37116，更多权益物料id敬请期待！
	_promotionId int64
}

// New
