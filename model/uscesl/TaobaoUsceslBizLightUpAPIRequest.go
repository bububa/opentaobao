package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizLightUpAPIRequest
价签LED等点亮 API请求
taobao.uscesl.biz.light.up

价签LED等点亮 */
type TaobaoUsceslBizLightUpAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 商家编号
	_bizBrandKey string
	// 价签条码
	_eslBarCode string
	// 亮灯颜色，绿：值为2；红：值为4
	_ledColor string
	// 亮灯时长，单位：秒，最大长度3600秒
	_lightUpTime int64
}

// New
