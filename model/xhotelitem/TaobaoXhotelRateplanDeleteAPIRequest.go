package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateplanDeleteAPIRequest
价格计划rateplan删除 API请求
taobao.xhotel.rateplan.delete

酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用 */
type TaobaoXhotelRateplanDeleteAPIRequest struct {
	model.Params
	// ratepland标识
	_rpId int64
	// 系统商，一般不用填写，使用须申请
	_vendor string
	// 商家价格政策编码
	_rateplanCode string
}

// New
