package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectPhoneSubmit 提交楼盘电话
// alibaba.alihouse.newhome.project.phone.submit
//
// 提交楼盘电话
func AlibabaAlihouseNewhomeProjectPhoneSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
