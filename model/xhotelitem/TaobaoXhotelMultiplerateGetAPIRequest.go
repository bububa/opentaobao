package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelMultiplerateGetAPIRequest
复杂房价查询接口 API请求
taobao.xhotel.multiplerate.get

查询复杂房价，支持通过入住人数，连住天数，商品信息，房价信息查询 */
type TaobaoXhotelMultiplerateGetAPIRequest struct {
	model.Params
	// 连住天数(范围1~10)
	_nod int64
	// 入住人数(范围1~10)
	_nop int64
	// 卖家的房价code
	_ratePlanCode string
	// 废弃，使用rate_plan_code
	_ratePlanId int64
	// 卖家的房型code
	_outRid string
	// 废弃，使用out_rid
	_gid int64
	// 系统商，一般不填写，使用须申请
	_vendor string
}

// New
