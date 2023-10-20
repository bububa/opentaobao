package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousestorepunishinfoquery 门店处罚信息查询
// alibaba.alihouse.store.punish.info.query
//
// 门店处罚信息查询
func Alibabaalihousestorepunishinfoquery(clt *core.SDKClient, req *alihouse.AlibabaalihousestorepunishinfoqueryAPIRequest, session string) (*alihouse.AlibabaalihousestorepunishinfoqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihousestorepunishinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
