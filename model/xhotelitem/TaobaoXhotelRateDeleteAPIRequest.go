package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateDeleteAPIRequest
rate删除接口 API请求
taobao.xhotel.rate.delete

酒店产品库rate删除 */
type TaobaoXhotelRateDeleteAPIRequest struct {
	model.Params
	// 系统商，一般不用填写，使用须申请
	_vendor string
	// 商家价格政策编码
	_rateplanCode string
	// 商家房型ID
	_outRid string
}

// New
