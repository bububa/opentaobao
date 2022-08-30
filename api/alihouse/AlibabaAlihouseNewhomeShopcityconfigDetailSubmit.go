package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopcityconfigDetailSubmit 城市配置信息接口
// alibaba.alihouse.newhome.shopcityconfig.detail.submit
//
// 上传城市配置信息
func AlibabaAlihouseNewhomeShopcityconfigDetailSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
