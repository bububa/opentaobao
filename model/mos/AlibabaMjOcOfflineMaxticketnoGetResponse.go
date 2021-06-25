package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
pos机获取线下最大小票号 APIResponse
alibaba.mj.oc.offline.maxticketno.get

给pos机提供线下最大小票号查询
*/
type AlibabaMjOcOfflineMaxticketnoGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjOcOfflineMaxticketnoGetResponse `json:"alibaba_mj_oc_offline_maxticketno_get_response,omitempty"`
}

type AlibabaMjOcOfflineMaxticketnoGetResponse struct {

    // 时间
    PayDate   string `json:"pay_date,omitempty"`

    // 收银机号
    PosNo   string `json:"pos_no,omitempty"`

    // 联合收银标记
    Union   bool `json:"union,omitempty"`

    // 小票号
    PayNo   string `json:"pay_no,omitempty"`

    // 外部门店号
    StoreNo   string `json:"store_no,omitempty"`

}
