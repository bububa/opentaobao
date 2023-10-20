package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxAuditCreativeModify 创意修改接口
// taobao.tanx.audit.creative.modify
//
// 创意修改接口
func TaobaoTanxAuditCreativeModify(clt *core.SDKClient, req *tanx.TaobaoTanxAuditCreativeModifyAPIRequest, resp *tanx.TaobaoTanxAuditCreativeModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
