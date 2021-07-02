package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// AlibabaMosFalconPosCounterQuery 云POS查看专柜属性
// alibaba.mos.falcon.pos.counter.query
//
// 银泰商业获取专柜是否支持小数等属性查看
func AlibabaMosFalconPosCounterQuery(clt *core.SDKClient, req *util.AlibabaMosFalconPosCounterQueryAPIRequest, session string) (*util.AlibabaMosFalconPosCounterQueryAPIResponse, error) {
	var resp util.AlibabaMosFalconPosCounterQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
