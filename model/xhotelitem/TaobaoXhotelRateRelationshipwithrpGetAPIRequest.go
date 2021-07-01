package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateRelationshipwithrpGetAPIRequest
根据gid查询卖家下所有的rpId API请求
taobao.xhotel.rate.relationshipwithrp.get

根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据 */
type TaobaoXhotelRateRelationshipwithrpGetAPIRequest struct {
	model.Params
	// 宝贝的gid
	_gid int64
	// 页数，可根据此值展示某页的数据。不填默认为1
	_pageNo int64
}

// New
