package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

/* AlibabaEleEnterpriseOrdernewGetstatus
订单状态查询接口
alibaba.ele.enterprise.ordernew.getstatus

订单状态查询接口 */
func AlibabaEleEnterpriseOrdernewGetstatus(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetstatusAPIRequest, session string) (*eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetstatusAPIResponse, error) {
	var resp eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
