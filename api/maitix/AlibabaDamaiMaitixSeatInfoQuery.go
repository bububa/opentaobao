package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixSeatInfoQuery 分销商查询座位信息
// alibaba.damai.maitix.seat.info.query
//
// 分销查询座位文案信息
func AlibabaDamaiMaitixSeatInfoQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixSeatInfoQueryAPIRequest, session string) (*maitix.AlibabaDamaiMaitixSeatInfoQueryAPIResponse, error) {
	var resp maitix.AlibabaDamaiMaitixSeatInfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
