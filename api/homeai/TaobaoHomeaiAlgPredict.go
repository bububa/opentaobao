package homeai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/homeai"
)

// Taobaohomeaialgpredict 硬装预测接口
// taobao.homeai.alg.predict
//
// 居然之家硬装预测服务
func Taobaohomeaialgpredict(clt *core.SDKClient, req *homeai.TaobaohomeaialgpredictAPIRequest, session string) (*homeai.TaobaohomeaialgpredictAPIResponse, error) {
	var resp homeai.TaobaohomeaialgpredictAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
