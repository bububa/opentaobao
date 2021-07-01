package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelInfoListGetForHelloAPIRequest
哈罗获取酒店详情 API请求
taobao.xhotel.info.list.get.for.hello

哈罗合作项目，供哈罗合作方批量和增量两种场景下查询已开通城市下的标准酒店及房型信息，不涉及用户登录信息 */
type TaobaoXhotelInfoListGetForHelloAPIRequest struct {
	model.Params
	// 参数封装模型
	_hotelInfoParam *HotelInfoParam
}

// New
