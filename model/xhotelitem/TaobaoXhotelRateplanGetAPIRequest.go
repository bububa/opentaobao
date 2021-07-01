package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateplanGetAPIRequest
价格计划rateplan查询 API请求
taobao.xhotel.rateplan.get

酒店产品库rateplan查询 */
type TaobaoXhotelRateplanGetAPIRequest struct {
	model.Params
	// 废弃，使用rateplan_code
	_rpid int64
	// 卖家自己系统的Code，简称RateCode
	_rateplanCode string
	// 系统商，一般不填写，使用须申请
	_vendor string
}

// New
