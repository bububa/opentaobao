package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// Alibabaalihealthpregnancynavigateinfoquery 查询底导数据
// alibaba.alihealth.pregnancy.navigateinfo.query
//
// 备孕管理--获取底部导航信息
func Alibabaalihealthpregnancynavigateinfoquery(clt *core.SDKClient, req *alihealthcrm.AlibabaalihealthpregnancynavigateinfoqueryAPIRequest, session string) (*alihealthcrm.AlibabaalihealthpregnancynavigateinfoqueryAPIResponse, error) {
	var resp alihealthcrm.AlibabaalihealthpregnancynavigateinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
