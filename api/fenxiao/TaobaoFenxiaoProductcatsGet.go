package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductcatsget 查询产品线列表
// taobao.fenxiao.productcats.get
//
// 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
func Taobaofenxiaoproductcatsget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductcatsgetAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductcatsgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductcatsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
