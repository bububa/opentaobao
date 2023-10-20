package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticstraceget 用户根据交易号查询物流流转信息
// taobao.logistics.trace.get
//
// 用户根据交易号查询物流流转信息
func Taobaologisticstraceget(clt *core.SDKClient, req *tblogistics.TaobaologisticstracegetAPIRequest, session string) (*tblogistics.TaobaologisticstracegetAPIResponse, error) {
	var resp tblogistics.TaobaologisticstracegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
