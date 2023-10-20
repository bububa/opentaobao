package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaovassubscsearch 订购记录导出
// taobao.vas.subsc.search
//
// 用于ISV查询自己名下的应用及收费项目的订购记录
func Taobaovassubscsearch(clt *core.SDKClient, req *servicecenter.TaobaovassubscsearchAPIRequest, session string) (*servicecenter.TaobaovassubscsearchAPIResponse, error) {
	var resp servicecenter.TaobaovassubscsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
