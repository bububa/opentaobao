package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingVersionCommit 提交版本号
// alibaba.hm.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
func AlibabaHmMarketingVersionCommit(clt *core.SDKClient, req *wdk.AlibabaHmMarketingVersionCommitAPIRequest, session string) (*wdk.AlibabaHmMarketingVersionCommitAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingVersionCommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
