package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaokaolascitemadd 考拉货品新增接口
// taobao.kaola.scitem.add
//
// 考拉货品新增接口
func Taobaokaolascitemadd(clt *core.SDKClient, req *fenxiao.TaobaokaolascitemaddAPIRequest, session string) (*fenxiao.TaobaokaolascitemaddAPIResponse, error) {
	var resp fenxiao.TaobaokaolascitemaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
