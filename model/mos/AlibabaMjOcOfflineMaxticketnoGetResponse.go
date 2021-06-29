package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
pos机获取线下最大小票号 API返回值 
alibaba.mj.oc.offline.maxticketno.get

给pos机提供线下最大小票号查询
*/
type AlibabaMjOcOfflineMaxticketnoGetAPIResponse struct {
    model.CommonResponse
    AlibabaMjOcOfflineMaxticketnoGetResponse
}

// pos机获取线下最大小票号 成功返回结果
type AlibabaMjOcOfflineMaxticketnoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_oc_offline_maxticketno_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 时间
    PayDate   string `json:"pay_date,omitempty" xml:"pay_date,omitempty"`
    // 收银机号
    PosNo   string `json:"pos_no,omitempty" xml:"pos_no,omitempty"`
    // 联合收银标记
    Union   bool `json:"union,omitempty" xml:"union,omitempty"`
    // 小票号
    PayNo   string `json:"pay_no,omitempty" xml:"pay_no,omitempty"`
    // 外部门店号
    StoreNo   string `json:"store_no,omitempty" xml:"store_no,omitempty"`
}
