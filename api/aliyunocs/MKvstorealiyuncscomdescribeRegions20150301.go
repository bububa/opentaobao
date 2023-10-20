package aliyunocs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyunocs"
)

// MKvstorealiyuncscomdescribeRegions20150301 查看Region列表
// m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01
//
// 查看Region列表
func MKvstorealiyuncscomdescribeRegions20150301(clt *core.SDKClient, req *aliyunocs.MKvstorealiyuncscomdescribeRegions20150301APIRequest, session string) (*aliyunocs.MKvstorealiyuncscomdescribeRegions20150301APIResponse, error) {
	var resp aliyunocs.MKvstorealiyuncscomdescribeRegions20150301APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
