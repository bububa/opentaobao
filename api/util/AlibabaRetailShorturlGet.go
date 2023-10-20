package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Alibabaretailshorturlget 短链接获取
// alibaba.retail.shorturl.get
//
// 短链接获取
func Alibabaretailshorturlget(clt *core.SDKClient, req *util.AlibabaretailshorturlgetAPIRequest, session string) (*util.AlibabaretailshorturlgetAPIResponse, error) {
	var resp util.AlibabaretailshorturlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
