package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* TmallServicecenterWorkcardUpdatelogistics
更新物流进度
tmall.servicecenter.workcard.updatelogistics

提供给外部合作服务商的物流进度更改接口 */
func TmallServicecenterWorkcardUpdatelogistics(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardUpdatelogisticsAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardUpdatelogisticsAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardUpdatelogisticsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
