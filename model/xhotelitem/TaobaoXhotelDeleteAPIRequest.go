package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelDeleteAPIRequest
删除酒店接口 API请求
taobao.xhotel.delete

删除飞猪酒店数据接口 */
type TaobaoXhotelDeleteAPIRequest struct {
	model.Params
	// 酒店id，传参方式  hid   或者   outer_id+vendor
	_hid int64
	// 酒店vendor
	_vendor string
	// 酒店编码
	_outerId string
}

// New
