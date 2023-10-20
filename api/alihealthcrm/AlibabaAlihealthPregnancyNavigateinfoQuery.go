package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthPregnancyNavigateinfoQuery 查询底导数据
// alibaba.alihealth.pregnancy.navigateinfo.query
//
// 备孕管理--获取底部导航信息
func AlibabaAlihealthPregnancyNavigateinfoQuery(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest, resp *alihealthcrm.AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
