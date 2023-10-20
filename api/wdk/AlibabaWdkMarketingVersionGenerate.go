package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingversiongenerate 生成发布使用的版本号
// alibaba.wdk.marketing.version.generate
//
// 生成发布使用的版本号
func Alibabawdkmarketingversiongenerate(clt *core.SDKClient, req *wdk.AlibabawdkmarketingversiongenerateAPIRequest, session string) (*wdk.AlibabawdkmarketingversiongenerateAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingversiongenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
