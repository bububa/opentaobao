package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelquotaupdate 库存更新接口
// taobao.xhotel.quota.update
//
// 库存更新接口
func Taobaoxhotelquotaupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelquotaupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelquotaupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelquotaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
