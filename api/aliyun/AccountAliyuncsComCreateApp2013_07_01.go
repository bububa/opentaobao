package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountAliyuncsComCreateApp2013_07_01 给指定用户创建appkey
// account.aliyuncs.com.CreateApp.2013-07-01
//
// 为某个用户创建appkey
func AccountAliyuncsComCreateApp2013_07_01(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateApp2013_07_01APIRequest, session string) (*aliyun.AccountAliyuncsComCreateApp2013_07_01APIResponse, error) {
	var resp aliyun.AccountAliyuncsComCreateApp2013_07_01APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
