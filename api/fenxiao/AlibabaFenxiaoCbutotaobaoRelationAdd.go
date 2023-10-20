package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Alibabafenxiaocbutotaobaorelationadd 1688分销铺货到淘宝关系添加
// alibaba.fenxiao.cbutotaobao.relation.add
//
// 1688分销铺货到淘宝关系添加
func Alibabafenxiaocbutotaobaorelationadd(clt *core.SDKClient, req *fenxiao.AlibabafenxiaocbutotaobaorelationaddAPIRequest, session string) (*fenxiao.AlibabafenxiaocbutotaobaorelationaddAPIResponse, error) {
	var resp fenxiao.AlibabafenxiaocbutotaobaorelationaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
