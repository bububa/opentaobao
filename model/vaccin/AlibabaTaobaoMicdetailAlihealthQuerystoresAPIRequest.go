package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest
疫苗预约门店列表查询 API请求
alibaba.taobao.micdetail.alihealth.querystores

微信小程序疫苗预约门店列表查询 */
type AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest struct {
	model.Params
	// 0不查询库存1查询库存
	_invType int64
	// 包ID
	_packageId int64
	// 页大小
	_pageSize int64
	// 页码
	_pageNum int64
	// 经度
	_lon string
	// 维度
	_lat string
	// 商品ID
	_itemId string
	// 用户ID
	_userId int64
	// 地区ID
	_districtId int64
	// 城市ID
	_cityId int64
	// 省份ID
	_provinceId int64
}

// New
