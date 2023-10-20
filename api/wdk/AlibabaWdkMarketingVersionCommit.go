package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingversioncommit 提交版本号
// alibaba.wdk.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
func Alibabawdkmarketingversioncommit(clt *core.SDKClient, req *wdk.AlibabawdkmarketingversioncommitAPIRequest, session string) (*wdk.AlibabawdkmarketingversioncommitAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingversioncommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
