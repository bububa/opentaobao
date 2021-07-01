package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelCityGetAPIRequest
酒店城市数据获取接口 API请求
taobao.xhotel.city.get

引流API，对外提供酒店城市数据 */
type TaobaoXhotelCityGetAPIRequest struct {
	model.Params
	// 分页读取的开始下标,从0开始
	_start int64
	// 分页读取的城市个数，最小值为1，最大值为200
	_count int64
}

// New
