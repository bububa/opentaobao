package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCrowdCrowdFinddmpcrowd 查询达摩盘精选人群模板
// taobao.onebp.dkx.crowd.crowd.finddmpcrowd
//
// 查询达摩盘精选人群模板；使用方法为先查询出topic和对应的templateId（需要一一对应），然后将想使用的模板topic&amp;templateId组填入Add接口中的new_dmp_template_crowd结构中提交即可。
func TaobaoOnebpDkxCrowdCrowdFinddmpcrowd(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest, resp *scs.TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
