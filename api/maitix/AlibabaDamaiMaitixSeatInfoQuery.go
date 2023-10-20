package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixseatinfoquery 分销商查询座位信息
// alibaba.damai.maitix.seat.info.query
//
// 分销查询座位文案信息
func Alibabadamaimaitixseatinfoquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixseatinfoqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixseatinfoqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixseatinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
