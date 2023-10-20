package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeshopcityconfigdetailsubmit 城市配置信息接口
// alibaba.alihouse.newhome.shopcityconfig.detail.submit
//
// 上传城市配置信息
func Alibabaalihousenewhomeshopcityconfigdetailsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeshopcityconfigdetailsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeshopcityconfigdetailsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeshopcityconfigdetailsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
