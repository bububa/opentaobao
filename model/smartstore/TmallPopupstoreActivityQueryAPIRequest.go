package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPopupstoreActivityQueryAPIRequest
查询某段时间内的快闪活动列表 API请求
tmall.popupstore.activity.query

提供给ISV查询某一时间段内包含指定appKey的活动列表 */
type TmallPopupstoreActivityQueryAPIRequest struct {
	model.Params
	// 查询开始时间,yyyy-MM-dd
	_startDate string
	// 查询结束时间，yyyy-MM-dd
	_endDate string
}

// New
