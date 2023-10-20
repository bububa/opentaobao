package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarleasecontractdownload 天猫开新车租后合同下载
// tmall.car.lease.contractdownload
//
// 天猫开新车租后合同下载
func Tmallcarleasecontractdownload(clt *core.SDKClient, req *tmallcar.TmallcarleasecontractdownloadAPIRequest, session string) (*tmallcar.TmallcarleasecontractdownloadAPIResponse, error) {
	var resp tmallcar.TmallcarleasecontractdownloadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
