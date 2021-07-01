package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

/* AlibabaOnetouchLogisticsExpressAddressStreetList
四级地址库-街道
alibaba.onetouch.logistics.express.address.street.list

四级地址库-街道模糊查询 */
func AlibabaOnetouchLogisticsExpressAddressStreetList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest, session string) (*icbulogistics.AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse, error) {
	var resp icbulogistics.AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
