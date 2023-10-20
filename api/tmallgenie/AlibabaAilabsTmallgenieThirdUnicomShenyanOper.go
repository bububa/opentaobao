package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieThirdUnicomShenyanOper 联通神眼注册操作
// alibaba.ailabs.tmallgenie.third.unicom.shenyan.oper
//
// 联通神眼注册操作
func AlibabaAilabsTmallgenieThirdUnicomShenyanOper(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
