package aliyun

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// OdpsAliyuncsComDescribeOdpsService20141215 查询ODPS服务
// odps.aliyuncs.com.DescribeOdpsService.2014-12-15
//
// 查询ODPS服务
func OdpsAliyuncsComDescribeOdpsService20141215(ctx context.Context, clt *core.SDKClient, req *aliyun.OdpsAliyuncsComDescribeOdpsService20141215APIRequest, resp *aliyun.OdpsAliyuncsComDescribeOdpsService20141215APIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
