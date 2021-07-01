package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* TmallServicecenterWorkcardCollect
工单揽件
tmall.servicecenter.workcard.collect

服务工单揽件接口 */
func TmallServicecenterWorkcardCollect(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCollectAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardCollectAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardCollectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
