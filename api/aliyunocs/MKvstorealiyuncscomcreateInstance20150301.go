package aliyunocs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyunocs"
)

// MKvstoreAliyuncsComCreateInstance20150301 创建OCS实例
// m-kvstore.aliyuncs.com.CreateInstance.2015-03-01
//
// 创建OCS实例
func MKvstoreAliyuncsComCreateInstance20150301(clt *core.SDKClient, req *aliyunocs.MKvstoreAliyuncsComCreateInstance20150301APIRequest, resp *aliyunocs.MKvstoreAliyuncsComCreateInstance20150301APIResponse, session string) error {
	return clt.Post(req, resp, session)
}
