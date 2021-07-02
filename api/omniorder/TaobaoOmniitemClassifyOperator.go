package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniitemClassifyOperator 添加/修改分类
// taobao.omniitem.classify.operator
//
// 添加/修改分类
func TaobaoOmniitemClassifyOperator(clt *core.SDKClient, req *omniorder.TaobaoOmniitemClassifyOperatorAPIRequest, session string) (*omniorder.TaobaoOmniitemClassifyOperatorAPIResponse, error) {
	var resp omniorder.TaobaoOmniitemClassifyOperatorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
