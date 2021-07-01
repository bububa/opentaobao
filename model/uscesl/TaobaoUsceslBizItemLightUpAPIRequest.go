package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizItemLightUpAPIRequest
商品条码亮灯API API请求
taobao.uscesl.biz.item.light.up

亮灯API */
type TaobaoUsceslBizItemLightUpAPIRequest struct {
	model.Params
	// 商品条码
	_itemBarCode string
	// 亮灯颜色，绿：值为2；红：值为4
	_ledColor string
	// 亮灯时长，单位：秒，最大长度3600秒
	_lightUpTime int64
	// 门店编号
	_storeId int64
	// 商家编号
	_bizBrandKey string
}

// New
