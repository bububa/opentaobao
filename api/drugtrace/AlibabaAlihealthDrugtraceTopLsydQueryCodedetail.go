package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsydquerycodedetail 根据码查询码信息
// alibaba.alihealth.drugtrace.top.lsyd.query.codedetail
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
func Alibabaalihealthdrugtracetoplsydquerycodedetail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsydquerycodedetailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsydquerycodedetailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsydquerycodedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
