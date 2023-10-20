package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// Tmallpopupstoreactivityquery 查询某段时间内的快闪活动列表
// tmall.popupstore.activity.query
//
// 提供给ISV查询某一时间段内包含指定appKey的活动列表
func Tmallpopupstoreactivityquery(clt *core.SDKClient, req *smartstore.TmallpopupstoreactivityqueryAPIRequest, session string) (*smartstore.TmallpopupstoreactivityqueryAPIResponse, error) {
	var resp smartstore.TmallpopupstoreactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
