package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// Alibabawestcrmcustomerinfoget 会员信息查询接口
// alibaba.westcrm.customer.info.get
//
// 会员信息查询接口
func Alibabawestcrmcustomerinfoget(clt *core.SDKClient, req *westcrm.AlibabawestcrmcustomerinfogetAPIRequest, session string) (*westcrm.AlibabawestcrmcustomerinfogetAPIResponse, error) {
	var resp westcrm.AlibabawestcrmcustomerinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
