package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaFenxiaoCbutotaobaoRelationAdd 1688分销铺货到淘宝关系添加
// alibaba.fenxiao.cbutotaobao.relation.add
//
// 1688分销铺货到淘宝关系添加
func AlibabaFenxiaoCbutotaobaoRelationAdd(clt *core.SDKClient, req *fenxiao.AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest, resp *fenxiao.AlibabaFenxiaoCbutotaobaoRelationAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
