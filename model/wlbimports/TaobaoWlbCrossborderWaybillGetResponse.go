package wlbimports

import (
    "github.com/bububa/opentaobao/model"
)

/* 
集货商家pdf和云打印面单获取，pdf需要配置白名单 APIResponse
taobao.wlb.crossborder.waybill.get

【TOF】欧洲供应商PDF格式电子面单渲染下发
 需求链接：https://aone.alibaba-inc.com/req/21210808
*/
type TaobaoWlbCrossborderWaybillGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbCrossborderWaybillGetResponse `json:"taobao_wlb_crossborder_waybill_get_response,omitempty"`
}

type TaobaoWlbCrossborderWaybillGetResponse struct {

    // 系统自动生成
    Result   *TopResult `json:"result,omitempty"`

}
