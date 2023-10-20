package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskucategorydelete 商家类目删除接口
// alibaba.wdk.sku.category.delete
//
// 商家类目删除接口
func Alibabawdkskucategorydelete(clt *core.SDKClient, req *wdk.AlibabawdkskucategorydeleteAPIRequest, session string) (*wdk.AlibabawdkskucategorydeleteAPIResponse, error) {
	var resp wdk.AlibabawdkskucategorydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
