package homeai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoHomeaiAlgPredictAPIRequest
硬装预测接口 API请求
taobao.homeai.alg.predict

居然之家硬装预测服务 */
type TaobaoHomeaiAlgPredictAPIRequest struct {
	model.Params
	// 来源房间json
	_fromCase string
	// 我的家房间json
	_toCase string
}

// New
