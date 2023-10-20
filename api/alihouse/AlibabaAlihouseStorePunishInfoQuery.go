package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseStorePunishInfoQuery 门店处罚信息查询
// alibaba.alihouse.store.punish.info.query
//
// 门店处罚信息查询
func AlibabaAlihouseStorePunishInfoQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseStorePunishInfoQueryAPIRequest, resp *alihouse.AlibabaAlihouseStorePunishInfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
