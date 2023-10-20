package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahtorderhotelsyncconfig 同步配置信息
// alibaba.htorder.hotel.sync.config
//
// 同步配置信息
func Alibabahtorderhotelsyncconfig(clt *core.SDKClient, req *happytrip.AlibabahtorderhotelsyncconfigAPIRequest, session string) (*happytrip.AlibabahtorderhotelsyncconfigAPIResponse, error) {
	var resp happytrip.AlibabahtorderhotelsyncconfigAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
