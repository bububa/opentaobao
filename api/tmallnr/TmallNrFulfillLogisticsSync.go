package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrFulfillLogisticsSync 同城配物流信息回传
// tmall.nr.fulfill.logistics.sync
//
// 同城配业务物流信息回传，通过接口将物流信息同步给天猫
func TmallNrFulfillLogisticsSync(clt *core.SDKClient, req *tmallnr.TmallNrFulfillLogisticsSyncAPIRequest, session string) (*tmallnr.TmallNrFulfillLogisticsSyncAPIResponse, error) {
	var resp tmallnr.TmallNrFulfillLogisticsSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
