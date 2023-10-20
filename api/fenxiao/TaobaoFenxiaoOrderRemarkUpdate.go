package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoorderremarkupdate 修改采购单备注
// taobao.fenxiao.order.remark.update
//
// 供应商修改采购单备注
func Taobaofenxiaoorderremarkupdate(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoorderremarkupdateAPIRequest, session string) (*fenxiao.TaobaofenxiaoorderremarkupdateAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoorderremarkupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
