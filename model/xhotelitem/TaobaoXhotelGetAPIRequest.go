package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelGetAPIRequest
酒店查询接口 API请求
taobao.xhotel.get

酒店查询接口 */
type TaobaoXhotelGetAPIRequest struct {
	model.Params
	// 废弃，请使用outer_id
	_hid int64
	// 卖家系统中的酒店ID。
	_outerId string
	// 系统商，一般不用填写，使用须申请
	_vendor string
	// 是否需要在售状态(默认false不需要)
	_needSaleInfo bool
}

// New
