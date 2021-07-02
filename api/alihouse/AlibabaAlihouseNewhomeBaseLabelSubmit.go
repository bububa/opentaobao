package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeBaseLabelSubmit 提交标签库
// alibaba.alihouse.newhome.base.label.submit
//
// 提交标签库
func AlibabaAlihouseNewhomeBaseLabelSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
