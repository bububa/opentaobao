package tuike

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuike"
)

// Alibabatuikeofferzhitoken 生成阿里口令
// alibaba.tuike.offer.zhitoken
//
// 推荐链接生产吱口令
func Alibabatuikeofferzhitoken(clt *core.SDKClient, req *tuike.AlibabatuikeofferzhitokenAPIRequest, session string) (*tuike.AlibabatuikeofferzhitokenAPIResponse, error) {
	var resp tuike.AlibabatuikeofferzhitokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
