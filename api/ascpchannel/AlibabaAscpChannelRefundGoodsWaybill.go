package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelrefundgoodswaybill 淘外分销退货回传物流单号
// alibaba.ascp.channel.refund.goods.waybill
//
// 淘外分销退货回传物流单号
func Alibabaascpchannelrefundgoodswaybill(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelrefundgoodswaybillAPIRequest, session string) (*ascpchannel.AlibabaascpchannelrefundgoodswaybillAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelrefundgoodswaybillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
