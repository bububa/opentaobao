package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// TaobaoItempropvaluesGet 获取标准类目属性值
// taobao.itempropvalues.get
//
// 获取标准类目属性值
func TaobaoItempropvaluesGet(clt *core.SDKClient, req *category.TaobaoItempropvaluesGetAPIRequest, session string) (*category.TaobaoItempropvaluesGetAPIResponse, error) {
	var resp category.TaobaoItempropvaluesGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
