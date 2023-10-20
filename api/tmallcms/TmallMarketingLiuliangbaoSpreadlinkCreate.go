package tmallcms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcms"
)

// Tmallmarketingliuliangbaospreadlinkcreate 创建流量宝活动链接
// tmall.marketing.liuliangbao.spreadlink.create
//
// 通过源活动链接和商家ID，创建流量宝活动链接
func Tmallmarketingliuliangbaospreadlinkcreate(clt *core.SDKClient, req *tmallcms.TmallmarketingliuliangbaospreadlinkcreateAPIRequest, session string) (*tmallcms.TmallmarketingliuliangbaospreadlinkcreateAPIResponse, error) {
	var resp tmallcms.TmallmarketingliuliangbaospreadlinkcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
