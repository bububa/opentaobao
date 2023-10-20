package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabstmallgeniethirdunicomshenyanoper 联通神眼注册操作
// alibaba.ailabs.tmallgenie.third.unicom.shenyan.oper
//
// 联通神眼注册操作
func Alibabaailabstmallgeniethirdunicomshenyanoper(clt *core.SDKClient, req *tmallgenie.AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest, session string) (*tmallgenie.AlibabaailabstmallgeniethirdunicomshenyanoperAPIResponse, error) {
	var resp tmallgenie.AlibabaailabstmallgeniethirdunicomshenyanoperAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
