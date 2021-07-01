package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

/* TmallPopupstoreActivityDeviceQuery
根据活动id查询活动相关快闪店及设备信息
tmall.popupstore.activity.device.query

查询某一活动的deviceCode的部署情况 */
func TmallPopupstoreActivityDeviceQuery(clt *core.SDKClient, req *smartstore.TmallPopupstoreActivityDeviceQueryAPIRequest, session string) (*smartstore.TmallPopupstoreActivityDeviceQueryAPIResponse, error) {
	var resp smartstore.TmallPopupstoreActivityDeviceQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
