package baichuanctg

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuanctg"
)

// Alibababaichuanctgtoutiaocontent 微博输出头条数据
// alibaba.baichuan.ctg.toutiao.content
//
// 百川头条内容获取
func Alibababaichuanctgtoutiaocontent(clt *core.SDKClient, req *baichuanctg.AlibababaichuanctgtoutiaocontentAPIRequest, session string) (*baichuanctg.AlibababaichuanctgtoutiaocontentAPIResponse, error) {
	var resp baichuanctg.AlibababaichuanctgtoutiaocontentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
