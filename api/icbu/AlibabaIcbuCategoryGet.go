package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbucategoryget 商品发布类目获取
// alibaba.icbu.category.get
//
// 获取商品发布类目
func Alibabaicbucategoryget(clt *core.SDKClient, req *icbu.AlibabaicbucategorygetAPIRequest, session string) (*icbu.AlibabaicbucategorygetAPIResponse, error) {
	var resp icbu.AlibabaicbucategorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
