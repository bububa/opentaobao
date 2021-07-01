package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusCityGetAPIRequest
城市接口 API请求
taobao.bus.city.get

汽车票出发城市获取接口，获取所有出发城市 */
type TaobaoBusCityGetAPIRequest struct {
	model.Params
}

// New
