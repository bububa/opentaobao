package lstmarketing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstmarketing"
)

// Alibabalstmarketingquerybyorderid 根据订单查询营销信息
// alibaba.lst.marketing.querybyorderid
//
// 根据订单查询营销信息
func Alibabalstmarketingquerybyorderid(clt *core.SDKClient, req *lstmarketing.AlibabalstmarketingquerybyorderidAPIRequest, session string) (*lstmarketing.AlibabalstmarketingquerybyorderidAPIResponse, error) {
	var resp lstmarketing.AlibabalstmarketingquerybyorderidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
