package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// Alibabaeleenterpriseordernewgetrefundinfo 退单和申诉
// alibaba.ele.enterprise.ordernew.getrefundinfo
//
// 退单和申诉
func Alibabaeleenterpriseordernewgetrefundinfo(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaeleenterpriseordernewgetrefundinfoAPIRequest, session string) (*eleenterpriseordernew.AlibabaeleenterpriseordernewgetrefundinfoAPIResponse, error) {
	var resp eleenterpriseordernew.AlibabaeleenterpriseordernewgetrefundinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
