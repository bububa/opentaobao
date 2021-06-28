package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增广告主接口 APIResponse
taobao.tanx.qualification.advertiser.add

外部dsp调用接口时会根据广告主名称和类型在tanx系统中新增一个广告主
*/
type TaobaoTanxQualificationAdvertiserAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tanx_qualification_advertiser_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"