package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductcatadd 新增产品线
// taobao.fenxiao.productcat.add
//
// 新增产品线
func Taobaofenxiaoproductcatadd(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductcataddAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductcataddAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductcataddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
