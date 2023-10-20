package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomelayoutsync 房通户型数据同步
// alibaba.alihouse.newhome.layout.sync
//
// 房通户型数据同步
func Alibabaalihousenewhomelayoutsync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomelayoutsyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomelayoutsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomelayoutsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
