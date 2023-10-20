package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelmultiplerateupdate 复杂价格推送接口（全量更新）
// taobao.xhotel.multiplerate.update
//
// 酒店产品库复杂rate（多人价，连住价等）更新
// 同时完全涵盖taobao.xhotel.rate.update的功能
func Taobaoxhotelmultiplerateupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelmultiplerateupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelmultiplerateupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelmultiplerateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
