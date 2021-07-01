package alitripbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripbp"
)

/* AlitripBpCouponinfoSync
飞猪广告券信息同步接口
alitrip.bp.couponinfo.sync

飞猪商业化券信息同步 */
func AlitripBpCouponinfoSync(clt *core.SDKClient, req *alitripbp.AlitripBpCouponinfoSyncAPIRequest, session string) (*alitripbp.AlitripBpCouponinfoSyncAPIResponse, error) {
	var resp alitripbp.AlitripBpCouponinfoSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
