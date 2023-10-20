package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelsubrefundcreate 淘外分销逆向创单（子单退）
// alibaba.ascp.channel.sub.refund.create
//
// 淘外分销逆向创单（子单退）
func Alibabaascpchannelsubrefundcreate(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelsubrefundcreateAPIRequest, session string) (*ascpchannel.AlibabaascpchannelsubrefundcreateAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelsubrefundcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
