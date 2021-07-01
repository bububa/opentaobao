package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

/* AlibabaMjOcSyncpayinfo
支付参考号回传
alibaba.mj.oc.syncpayinfo

支付参考号同步到oc */
func AlibabaMjOcSyncpayinfo(clt *core.SDKClient, req *mos.AlibabaMjOcSyncpayinfoAPIRequest, session string) (*mos.AlibabaMjOcSyncpayinfoAPIResponse, error) {
	var resp mos.AlibabaMjOcSyncpayinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
