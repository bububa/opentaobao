package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeitaoFeedSynchronizeAPIRequest
推广淘小铺isv 活动到微淘feed API请求
taobao.weitao.feed.synchronize

推广淘小铺isv 活动到微淘feed */
type TaobaoWeitaoFeedSynchronizeAPIRequest struct {
	model.Params
	// 活动id
	_bizId int64
	// feed封面图片url
	_coverPath string
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
}

// New
