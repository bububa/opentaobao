package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// Alibabaicburfqrecommend rfq推荐
// alibaba.icbu.rfq.recommend
//
// rfq推荐
func Alibabaicburfqrecommend(clt *core.SDKClient, req *icburfq.AlibabaicburfqrecommendAPIRequest, session string) (*icburfq.AlibabaicburfqrecommendAPIResponse, error) {
	var resp icburfq.AlibabaicburfqrecommendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
