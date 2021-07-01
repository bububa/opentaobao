package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizEslInfoGetAPIRequest
价签设备信息查询接口 API请求
taobao.uscesl.biz.esl.info.get

价签设备信息查询接口 */
type TaobaoUsceslBizEslInfoGetAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 门店storeId
	_storeId int64
	// 商家ID
	_bizBrandKey string
}

// New
