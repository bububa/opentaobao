package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// Alibabaicburfqmyequity 我的权益
// alibaba.icbu.rfq.myequity
//
// 查询供应商权益接口
func Alibabaicburfqmyequity(clt *core.SDKClient, req *icburfq.AlibabaicburfqmyequityAPIRequest, session string) (*icburfq.AlibabaicburfqmyequityAPIResponse, error) {
	var resp icburfq.AlibabaicburfqmyequityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
