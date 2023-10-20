package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopconfigDetailSubmit 店铺配置信息接口
// alibaba.alihouse.newhome.shopconfig.detail.submit
//
// 提供店铺配置的能力
func AlibabaAlihouseNewhomeShopconfigDetailSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
