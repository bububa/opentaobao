package alitripbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪广告券信息同步接口 API返回值 
alitrip.bp.couponinfo.sync

飞猪商业化券信息同步
*/
type AlitripBpCouponinfoSyncAPIResponse struct {
    model.CommonResponse
    AlitripBpCouponinfoSyncAPIResponseModel
}

// 飞猪广告券信息同步接口 成功返回结果
type AlitripBpCouponinfoSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_bp_couponinfo_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *AdResult `json:"result,omitempty" xml:"result,omitempty"`
}
