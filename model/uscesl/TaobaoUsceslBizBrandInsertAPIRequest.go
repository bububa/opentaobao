package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizBrandInsertAPIRequest
新增电子价签商家 API请求
taobao.uscesl.biz.brand.insert

一个电子价签业务身份下新增商家接口 */
type TaobaoUsceslBizBrandInsertAPIRequest struct {
	model.Params
	// 商家名称
	_brandName string
	// 商家外部编号
	_brandOutCode string
}

// New
