package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsDataSportsSyncstatdata 阿里体育数据中心用户当天累积数据同步接口
// alibaba.alisports.data.sports.syncstatdata
//
// 阿里体育数据中心用户当天累积数据同步接口
func AlibabaAlisportsDataSportsSyncstatdata(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsDataSportsSyncstatdataAPIRequest, resp *alisports.AlibabaAlisportsDataSportsSyncstatdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
