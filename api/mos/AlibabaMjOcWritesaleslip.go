package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

/* AlibabaMjOcWritesaleslip
开票占库
alibaba.mj.oc.writesaleslip

开票占库 */
func AlibabaMjOcWritesaleslip(clt *core.SDKClient, req *mos.AlibabaMjOcWritesaleslipAPIRequest, session string) (*mos.AlibabaMjOcWritesaleslipAPIResponse, error) {
	var resp mos.AlibabaMjOcWritesaleslipAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
