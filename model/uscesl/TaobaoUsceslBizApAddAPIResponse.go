package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增价签通讯AP设备 API返回值 
taobao.uscesl.biz.ap.add

根据门店和ap的MAC地址新增
*/
type TaobaoUsceslBizApAddAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizApAddAPIResponseModel
}

// 新增价签通讯AP设备 成功返回结果
type TaobaoUsceslBizApAddAPIResponseModel struct {
    XMLName xml.Name `xml:"uscesl_biz_ap_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoUsceslBizApAddResult `json:"result,omitempty" xml:"result,omitempty"`
}
