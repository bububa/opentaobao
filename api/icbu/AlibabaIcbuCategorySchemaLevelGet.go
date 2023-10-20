package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbucategoryschemalevelget (新)层级属性获取
// alibaba.icbu.category.schema.level.get
//
// 将表单中层级属性的子属性返回
func Alibabaicbucategoryschemalevelget(clt *core.SDKClient, req *icbu.AlibabaicbucategoryschemalevelgetAPIRequest, session string) (*icbu.AlibabaicbucategoryschemalevelgetAPIResponse, error) {
	var resp icbu.AlibabaicbucategoryschemalevelgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
