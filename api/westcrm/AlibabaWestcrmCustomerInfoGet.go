package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

/* AlibabaWestcrmCustomerInfoGet
会员信息查询接口
alibaba.westcrm.customer.info.get

会员信息查询接口 */
func AlibabaWestcrmCustomerInfoGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmCustomerInfoGetAPIRequest, session string) (*westcrm.AlibabaWestcrmCustomerInfoGetAPIResponse, error) {
	var resp westcrm.AlibabaWestcrmCustomerInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
