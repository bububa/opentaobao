package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomelinesync 环线数据同步
// alibaba.alihouse.newhome.line.sync
//
// 环线数据同步
func Alibabaalihousenewhomelinesync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomelinesyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomelinesyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomelinesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
