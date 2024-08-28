package smartstore

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// TmallPopupstoreActivityDeviceQuery 根据活动id查询活动相关快闪店及设备信息
// tmall.popupstore.activity.device.query
//
// 查询某一活动的deviceCode的部署情况
func TmallPopupstoreActivityDeviceQuery(ctx context.Context, clt *core.SDKClient, req *smartstore.TmallPopupstoreActivityDeviceQueryAPIRequest, resp *smartstore.TmallPopupstoreActivityDeviceQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
