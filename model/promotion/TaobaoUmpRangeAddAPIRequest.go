package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpRangeAddAPIRequest
添加活动范围 API请求
taobao.ump.range.add

指定某项活动中，某个商家的某些类型物品（指定商品或者别的）参加或者不参加活动。当活动详情的参与类型为部分或者部分不参加的时候，需要指定具体哪部分参加或者不参加，使用本接口完成操作。比如部分商品满就送，这里的range用来指定具体哪些商品参加满就送活动。 */
type TaobaoUmpRangeAddAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
	// 范围的类型，比如是全店，商品，见：MarketingConstants.PARTICIPATE_TYPE_*
	_type int64
	// id列表，当范围类型为商品时，该id为商品id.多个id用逗号隔开，一次不超过50个
	_ids string
}

// New
