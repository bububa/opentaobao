package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRoomGetAPIRequest
room实体查询 API请求
taobao.xhotel.room.get

此接口用于查询一个商品，根据传入的gid查询商品信息。卖家只能查询自己的商品。 */
type TaobaoXhotelRoomGetAPIRequest struct {
	model.Params
	// 卖家渠道 如果gid为空，那么out_rid和vendor都不能为空。 支持通过gid或者通过out_rid和vendor来获取商品
	_vendor string
	// 外部房型id 如果gid为空，那么out_rid和vendor都不能为空 支持通过gid或者通过out_rid和vendor来获取商品
	_outRid string
	// 废弃
	_gid int64
}

// New
