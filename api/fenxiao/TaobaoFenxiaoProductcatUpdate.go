package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductcatUpdate 修改产品线
// taobao.fenxiao.productcat.update
//
// 修改产品线
func TaobaoFenxiaoProductcatUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatUpdateAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductcatUpdateAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductcatUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
