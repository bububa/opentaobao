package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// Alibabaeleenterpriseordernewgetstatus 订单状态查询接口
// alibaba.ele.enterprise.ordernew.getstatus
//
// 订单状态查询接口
func Alibabaeleenterpriseordernewgetstatus(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaeleenterpriseordernewgetstatusAPIRequest, session string) (*eleenterpriseordernew.AlibabaeleenterpriseordernewgetstatusAPIResponse, error) {
	var resp eleenterpriseordernew.AlibabaeleenterpriseordernewgetstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
