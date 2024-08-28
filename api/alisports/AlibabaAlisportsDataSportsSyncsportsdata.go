package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsDataSportsSyncsportsdata 阿里体育数据中心用户运动数据同步接口
// alibaba.alisports.data.sports.syncsportsdata
//
// 阿里体育数据中心用户运动数据同步接口
func AlibabaAlisportsDataSportsSyncsportsdata(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsDataSportsSyncsportsdataAPIRequest, resp *alisports.AlibabaAlisportsDataSportsSyncsportsdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
