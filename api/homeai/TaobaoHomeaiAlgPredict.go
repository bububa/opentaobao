package homeai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/homeai"
)

// TaobaoHomeaiAlgPredict 硬装预测接口
// taobao.homeai.alg.predict
//
// 居然之家硬装预测服务
func TaobaoHomeaiAlgPredict(clt *core.SDKClient, req *homeai.TaobaoHomeaiAlgPredictAPIRequest, resp *homeai.TaobaoHomeaiAlgPredictAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
