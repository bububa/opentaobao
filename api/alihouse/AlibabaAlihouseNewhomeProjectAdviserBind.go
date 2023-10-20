package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectAdviserBind 置业顾问批量绑定楼盘
// alibaba.alihouse.newhome.project.adviser.bind
//
// 置业顾问批量绑定楼盘
func AlibabaAlihouseNewhomeProjectAdviserBind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
