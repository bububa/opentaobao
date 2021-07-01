package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

/* OdpsAliyuncsComDescribeOdpsService2014_12_15
查询ODPS服务
odps.aliyuncs.com.DescribeOdpsService.2014-12-15

查询ODPS服务 */
func OdpsAliyuncsComDescribeOdpsService2014_12_15(clt *core.SDKClient, req *aliyun.OdpsAliyuncsComDescribeOdpsService2014_12_15APIRequest, session string) (*aliyun.OdpsAliyuncsComDescribeOdpsService2014_12_15APIResponse, error) {
	var resp aliyun.OdpsAliyuncsComDescribeOdpsService2014_12_15APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
