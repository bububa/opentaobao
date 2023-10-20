package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaooctradetracealertsgetAPIResponse 异常订单信息获取 API返回值
// taobao.oc.tradetrace.alerts.get
//
// 提供订单预警模块的数据查询接口
type TaobaooctradetracealertsgetAPIResponse struct {
	model.CommonResponse
	TaobaooctradetracealertsgetAPIResponseModel
}

// TaobaooctradetracealertsgetAPIResponseModel is 异常订单信息获取 成功返回结果
type TaobaooctradetracealertsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"oc_tradetrace_alerts_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异常订单数据
	ResultList []SimpleAbnormalOrderDetail `json:"result_list,omitempty" xml:"result_list>simple_abnormal_order_detail,omitempty"`
}
