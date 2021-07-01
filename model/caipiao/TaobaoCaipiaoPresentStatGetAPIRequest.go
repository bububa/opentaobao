package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCaipiaoPresentStatGetAPIRequest
获取卖家按天统计的彩票赠送数据 API请求
taobao.caipiao.present.stat.get

查询卖家一段时间内按天统计的彩票赠送数据，只支持查询90天以内的数据. */
type TaobaoCaipiaoPresentStatGetAPIRequest struct {
	model.Params
	// 指定查询的天数，从当前日期（不包括当前日期）向前推算的天数，可为空。如果为空、0、负数或者大于90天，则设置为默认的90天。举例：当天是20120703， days=2， 则统计数据的日期为：20120702，20120701.
	_days int64
}

// New
