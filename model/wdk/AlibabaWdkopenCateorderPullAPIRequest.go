package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkopenCateorderPullAPIRequest
商户回传餐饮加工单状态 API请求
alibaba.wdkopen.cateorder.pull

商户回传餐饮加工单状态 */
type AlibabaWdkopenCateorderPullAPIRequest struct {
	model.Params
	// 经营店ID
	_storeId string
	// 回传状态,PREPARING,准备中，制作中；PRODUCE_FINISH，制作完成；FETCHED 已取餐；  CANCEL，加工失败/取消
	_status string
	// 主站主订单ID
	_outOrderId string
	// 主站子订单ID列表, 为空则表示回传整单状态
	_subOutOrderIds []string
}

// New
