package aliyunocs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyunocs"
)

// MKvstorealiyuncscomcreateInstance20150301 创建OCS实例
// m-kvstore.aliyuncs.com.CreateInstance.2015-03-01
//
// 创建OCS实例
func MKvstorealiyuncscomcreateInstance20150301(clt *core.SDKClient, req *aliyunocs.MKvstorealiyuncscomcreateInstance20150301APIRequest, session string) (*aliyunocs.MKvstorealiyuncscomcreateInstance20150301APIResponse, error) {
	var resp aliyunocs.MKvstorealiyuncscomcreateInstance20150301APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
