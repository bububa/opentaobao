package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbucategorygetnew （新）ICBU类目树获取接口
// alibaba.icbu.category.get.new
//
// 获取商品发布类目
func Alibabaicbucategorygetnew(clt *core.SDKClient, req *icbu.AlibabaicbucategorygetnewAPIRequest, session string) (*icbu.AlibabaicbucategorygetnewAPIResponse, error) {
	var resp icbu.AlibabaicbucategorygetnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
