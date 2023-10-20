package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingversioncommit 提交版本号
// alibaba.hm.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
func Alibabahmmarketingversioncommit(clt *core.SDKClient, req *wdk.AlibabahmmarketingversioncommitAPIRequest, session string) (*wdk.AlibabahmmarketingversioncommitAPIResponse, error) {
	var resp wdk.AlibabahmmarketingversioncommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
