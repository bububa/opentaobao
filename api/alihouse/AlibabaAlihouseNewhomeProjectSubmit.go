package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectSubmit 提交楼盘信息
// alibaba.alihouse.newhome.project.submit
//
// 提交楼盘信息
func AlibabaAlihouseNewhomeProjectSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
