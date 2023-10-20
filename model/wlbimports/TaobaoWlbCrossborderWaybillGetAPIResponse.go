package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbcrossborderwaybillgetAPIResponse 集货商家pdf和云打印面单获取，pdf需要配置白名单 API返回值
// taobao.wlb.crossborder.waybill.get
//
// 【TOF】欧洲供应商PDF格式电子面单渲染下发
//
//	需求链接：https://aone.alibaba-inc.com/req/21210808
type TaobaowlbcrossborderwaybillgetAPIResponse struct {
	model.CommonResponse
	TaobaowlbcrossborderwaybillgetAPIResponseModel
}

// TaobaowlbcrossborderwaybillgetAPIResponseModel is 集货商家pdf和云打印面单获取，pdf需要配置白名单 成功返回结果
type TaobaowlbcrossborderwaybillgetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_crossborder_waybill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *TaobaowlbcrossborderwaybillgetTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
