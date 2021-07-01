package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeitaoFeedSynchronizeNewAPIRequest
推广淘小铺isv 活动到微淘feed API请求
taobao.weitao.feed.synchronize.new

推广微淘互动应用活动到微淘 */
type TaobaoWeitaoFeedSynchronizeNewAPIRequest struct {
	model.Params
	// 广播类型：粉丝猜价格461、投票有礼462、粉丝抢红包463、盖楼有礼464、加购有礼465
	_feedType int64
	// feed点击跳转的活动地址
	_detailUrl string
	// feed展示结束时间
	_endTime int64
	// feed展示开始时间
	_startTime int64
	// feed描述
	_summary string
	// feed标题
	_title string
	// 宝贝列表，用于card展示，0~2个宝贝ID
	_itemIds []string
	// 活动ID
	_sbizId string
}

// New
