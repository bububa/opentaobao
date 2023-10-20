package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpresslogisticsdstrackinginfoquery 查询物流追踪信息
// aliexpress.logistics.ds.trackinginfo.query
//
// Dropshipper查询物流追踪信息
func Aliexpresslogisticsdstrackinginfoquery(clt *core.SDKClient, req *aedropshiper.AliexpresslogisticsdstrackinginfoqueryAPIRequest, session string) (*aedropshiper.AliexpresslogisticsdstrackinginfoqueryAPIResponse, error) {
	var resp aedropshiper.AliexpresslogisticsdstrackinginfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
