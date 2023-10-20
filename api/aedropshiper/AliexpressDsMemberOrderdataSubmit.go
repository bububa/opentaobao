package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsMemberOrderdataSubmit dropshipper数据回流
// aliexpress.ds.member.orderdata.submit
//
// dropshipper数据回流
func AliexpressDsMemberOrderdataSubmit(clt *core.SDKClient, req *aedropshiper.AliexpressDsMemberOrderdataSubmitAPIRequest, session string) (*aedropshiper.AliexpressDsMemberOrderdataSubmitAPIResponse, error) {
	var resp aedropshiper.AliexpressDsMemberOrderdataSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
