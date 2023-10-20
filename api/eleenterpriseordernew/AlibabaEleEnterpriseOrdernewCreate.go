package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// Alibabaeleenterpriseordernewcreate 创建订单
// alibaba.ele.enterprise.ordernew.create
//
// 创建订单
func Alibabaeleenterpriseordernewcreate(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaeleenterpriseordernewcreateAPIRequest, session string) (*eleenterpriseordernew.AlibabaeleenterpriseordernewcreateAPIResponse, error) {
	var resp eleenterpriseordernew.AlibabaeleenterpriseordernewcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
