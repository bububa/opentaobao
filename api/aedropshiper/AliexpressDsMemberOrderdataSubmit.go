package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsMemberOrderdataSubmit dropshipper数据回流
// aliexpress.ds.member.orderdata.submit
//
// dropshipper数据回流
func AliexpressDsMemberOrderdataSubmit(clt *core.SDKClient, req *aedropshiper.AliexpressDsMemberOrderdataSubmitAPIRequest, resp *aedropshiper.AliexpressDsMemberOrderdataSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
