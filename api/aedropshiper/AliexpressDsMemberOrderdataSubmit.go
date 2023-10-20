package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpressdsmemberorderdatasubmit dropshipper数据回流
// aliexpress.ds.member.orderdata.submit
//
// dropshipper数据回流
func Aliexpressdsmemberorderdatasubmit(clt *core.SDKClient, req *aedropshiper.AliexpressdsmemberorderdatasubmitAPIRequest, session string) (*aedropshiper.AliexpressdsmemberorderdatasubmitAPIResponse, error) {
	var resp aedropshiper.AliexpressdsmemberorderdatasubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
