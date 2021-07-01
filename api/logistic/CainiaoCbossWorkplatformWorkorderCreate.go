package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* CainiaoCbossWorkplatformWorkorderCreate
菜鸟工单创建接口
cainiao.cboss.workplatform.workorder.create

菜鸟工单创建接口，目前调用者ISV */
func CainiaoCbossWorkplatformWorkorderCreate(clt *core.SDKClient, req *logistic.CainiaoCbossWorkplatformWorkorderCreateAPIRequest, session string) (*logistic.CainiaoCbossWorkplatformWorkorderCreateAPIResponse, error) {
	var resp logistic.CainiaoCbossWorkplatformWorkorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
