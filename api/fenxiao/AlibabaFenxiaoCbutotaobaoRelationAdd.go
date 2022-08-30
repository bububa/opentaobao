package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaFenxiaoCbutotaobaoRelationAdd 1688分销铺货到淘宝关系添加
// alibaba.fenxiao.cbutotaobao.relation.add
//
// 1688分销铺货到淘宝关系添加
func AlibabaFenxiaoCbutotaobaoRelationAdd(clt *core.SDKClient, req *fenxiao.AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest, session string) (*fenxiao.AlibabaFenxiaoCbutotaobaoRelationAddAPIResponse, error) {
	var resp fenxiao.AlibabaFenxiaoCbutotaobaoRelationAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
