package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// Tmallpopupstoreactivitydevicequery 根据活动id查询活动相关快闪店及设备信息
// tmall.popupstore.activity.device.query
//
// 查询某一活动的deviceCode的部署情况
func Tmallpopupstoreactivitydevicequery(clt *core.SDKClient, req *smartstore.TmallpopupstoreactivitydevicequeryAPIRequest, session string) (*smartstore.TmallpopupstoreactivitydevicequeryAPIResponse, error) {
	var resp smartstore.TmallpopupstoreactivitydevicequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
