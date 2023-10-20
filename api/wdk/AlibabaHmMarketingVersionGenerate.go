package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingversiongenerate 生成发布使用的版本号
// alibaba.hm.marketing.version.generate
//
// 生成发布使用的版本号
func Alibabahmmarketingversiongenerate(clt *core.SDKClient, req *wdk.AlibabahmmarketingversiongenerateAPIRequest, session string) (*wdk.AlibabahmmarketingversiongenerateAPIResponse, error) {
	var resp wdk.AlibabahmmarketingversiongenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
