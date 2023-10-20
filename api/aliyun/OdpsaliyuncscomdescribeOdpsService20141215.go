package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// OdpsAliyuncsComDescribeOdpsService20141215 查询ODPS服务
// odps.aliyuncs.com.DescribeOdpsService.2014-12-15
//
// 查询ODPS服务
func OdpsAliyuncsComDescribeOdpsService20141215(clt *core.SDKClient, req *aliyun.OdpsAliyuncsComDescribeOdpsService20141215APIRequest, session string) (*aliyun.OdpsAliyuncsComDescribeOdpsService20141215APIResponse, error) {
	var resp aliyun.OdpsAliyuncsComDescribeOdpsService20141215APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
