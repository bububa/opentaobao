package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoNewretailDivisionRecordListGet 导购分佣明细列表
// taobao.newretail.division.record.list.get
//
// 提供分页查询导购分佣明细的能力
func TaobaoNewretailDivisionRecordListGet(clt *core.SDKClient, req *user.TaobaoNewretailDivisionRecordListGetAPIRequest, resp *user.TaobaoNewretailDivisionRecordListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
