package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsDataSportsSyncsleepdata 阿里体育数据中心用户睡眠数据同步接口
// alibaba.alisports.data.sports.syncsleepdata
//
// 阿里体育数据中心用户睡眠数据同步接口
func AlibabaAlisportsDataSportsSyncsleepdata(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsDataSportsSyncsleepdataAPIRequest, resp *alisports.AlibabaAlisportsDataSportsSyncsleepdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
