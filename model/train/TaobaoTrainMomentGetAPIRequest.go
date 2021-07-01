package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainMomentGetAPIRequest
火车票时刻表 API请求
taobao.train.moment.get

查询火车票车次时刻表 */
type TaobaoTrainMomentGetAPIRequest struct {
	model.Params
	// 出参
	_param *TrainMomentTopParam
}

// New
