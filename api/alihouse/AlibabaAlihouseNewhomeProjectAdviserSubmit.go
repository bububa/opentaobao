package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectAdviserSubmit 提交楼盘顾问
// alibaba.alihouse.newhome.project.adviser.submit
//
// 提交楼盘顾问
func AlibabaAlihouseNewhomeProjectAdviserSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectAdviserSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
