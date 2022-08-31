package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivitySave 新增或者更新销售活动
// alibaba.alihouse.newhome.activity.save
//
// 新增或者更新销售活动
func AlibabaAlihouseNewhomeActivitySave(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivitySaveAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeActivitySaveAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeActivitySaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
