package aliyunocs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyunocs"
)

// MKvstoreAliyuncsComDescribeRegions20150301 查看Region列表
// m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01
//
// 查看Region列表
func MKvstoreAliyuncsComDescribeRegions20150301(ctx context.Context, clt *core.SDKClient, req *aliyunocs.MKvstoreAliyuncsComDescribeRegions20150301APIRequest, resp *aliyunocs.MKvstoreAliyuncsComDescribeRegions20150301APIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
