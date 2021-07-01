package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmActivityGetAPIRequest
私域导购查询活动详情 API请求
alibaba.lsy.crm.activity.get

私域导购查询活动详情 */
type AlibabaLsyCrmActivityGetAPIRequest struct {
	model.Params
	// 活动id
	_activityId int64
	// 导购员id
	_guiderId int64
	// 摊位id
	_storeId int64
}

// New
