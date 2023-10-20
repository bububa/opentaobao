package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectSubmit 提交楼盘信息
// alibaba.alihouse.newhome.project.submit
//
// 提交楼盘信息
func AlibabaAlihouseNewhomeProjectSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
