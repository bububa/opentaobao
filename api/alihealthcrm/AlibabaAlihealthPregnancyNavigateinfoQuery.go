package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

/* AlibabaAlihealthPregnancyNavigateinfoQuery
查询底导数据
alibaba.alihealth.pregnancy.navigateinfo.query

备孕管理--获取底部导航信息 */
func AlibabaAlihealthPregnancyNavigateinfoQuery(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse, error) {
	var resp alihealthcrm.AlibabaAlihealthPregnancyNavigateinfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
