package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// AlibabaMosFalconPosCounterQuery 云POS查看专柜属性
// alibaba.mos.falcon.pos.counter.query
//
// 银泰商业获取专柜是否支持小数等属性查看
func AlibabaMosFalconPosCounterQuery(ctx context.Context, clt *core.SDKClient, req *util.AlibabaMosFalconPosCounterQueryAPIRequest, resp *util.AlibabaMosFalconPosCounterQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
