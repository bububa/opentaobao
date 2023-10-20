package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelrefundclose 渠道退款单关闭
// alibaba.ascp.channel.refund.close
//
// 渠道退款单关闭
func Alibabaascpchannelrefundclose(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelrefundcloseAPIRequest, session string) (*ascpchannel.AlibabaascpchannelrefundcloseAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelrefundcloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
