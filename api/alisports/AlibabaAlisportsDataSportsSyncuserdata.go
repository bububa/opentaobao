package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsDataSportsSyncuserdata 阿里体育数据中心用户个人信息同步接口
// alibaba.alisports.data.sports.syncuserdata
//
// 阿里体育数据中心用户个人信息同步接口
func AlibabaAlisportsDataSportsSyncuserdata(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsDataSportsSyncuserdataAPIRequest, resp *alisports.AlibabaAlisportsDataSportsSyncuserdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
