package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductcatdelete 删除产品线
// taobao.fenxiao.productcat.delete
//
// 删除产品线
func Taobaofenxiaoproductcatdelete(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductcatdeleteAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductcatdeleteAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductcatdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
