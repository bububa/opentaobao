package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizEslBindAPIRequest
电子价签绑定接口 API请求
taobao.uscesl.biz.esl.bind

电子价签商品绑定接口 */
type TaobaoUsceslBizEslBindAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 商品条码
	_itemBarCode string
	// 门店storeId
	_storeId int64
	// 商家ID
	_bizBrandKey string
	// 额外扩展信息
	_extendInfo string
}

// New
