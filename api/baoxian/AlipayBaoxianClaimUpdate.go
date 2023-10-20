package baoxian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// Alipaybaoxianclaimupdate 更新赔案
// alipay.baoxian.claim.update
//
// 更新保险理赔单
func Alipaybaoxianclaimupdate(clt *core.SDKClient, req *baoxian.AlipaybaoxianclaimupdateAPIRequest, session string) (*baoxian.AlipaybaoxianclaimupdateAPIResponse, error) {
	var resp baoxian.AlipaybaoxianclaimupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
