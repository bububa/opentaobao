package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseCategoryControlSync 类目权限上翻
// alibaba.alihouse.category.control.sync
//
// 类目权限上翻
func AlibabaAlihouseCategoryControlSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseCategoryControlSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseCategoryControlSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseCategoryControlSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
