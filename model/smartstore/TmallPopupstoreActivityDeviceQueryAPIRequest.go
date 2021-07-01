package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPopupstoreActivityDeviceQueryAPIRequest
根据活动id查询活动相关快闪店及设备信息 API请求
tmall.popupstore.activity.device.query

查询某一活动的deviceCode的部署情况 */
type TmallPopupstoreActivityDeviceQueryAPIRequest struct {
	model.Params
	// ISV的活动ID
	_activityId int64
}

// New
